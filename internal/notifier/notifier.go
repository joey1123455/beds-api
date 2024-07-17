package notifier

import (
	"fmt"
	"sync"
)

type Notifier interface {
	SendNotification(templateName string, data *Notification) error
	Key() string
}

type NotificationSender struct {
	notifiers map[string]Notifier
	mu        sync.Mutex
}

func NewNotificationSender(notifiers ...Notifier) *NotificationSender {
	notifyMap := make(map[string]Notifier)
	for _, n := range notifiers {
		notifyMap[n.Key()] = n
	}
	return &NotificationSender{notifiers: notifyMap}
}

func (ns *NotificationSender) SendOnChannel(channel string, notification *Notification) error {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	if notifier, exists := ns.notifiers[channel]; exists {
		return notifier.SendNotification(notification.TemplateName, notification)
	}
	return fmt.Errorf("channel %s not found", channel)
}

func (ns *NotificationSender) SendNotification(notification *Notification) error {
	for _, channel := range notification.GetChannels() {
		if notifier, ok := ns.notifiers[channel]; ok {
			r := notification.Recipient
			switch {
			case r.GetEmail() != "" && r.GetPhone() != "" && (channel == NotificationChannelEmail || channel == NotificationChannelSMS):

				err := notifier.SendNotification(notification.TemplateName, notification)
				if err != nil {
					return err
				}
			case r.GetEmail() != "" && channel == NotificationChannelEmail:

				err := notifier.SendNotification(notification.TemplateName, notification)
				if err != nil {
					return err
				}

			case r.GetPhone() != "" && channel == NotificationChannelSMS:
				err := notifier.SendNotification(notification.TemplateName, notification)
				if err != nil {
					return err
				}

			default:
				return fmt.Errorf("invalid recipient: %+v", r)
			}

		}
	}
	return nil
}
