package notifier

import "sync"

const (
	NotificationChannelEmail = "email"
	NotificationChannelSMS   = "sms"
)

// Notification data structure
type Notification struct {
	Recipient    Recipient              `json:"recipient"`
	Title        string                 `json:"title"`
	TemplateName string                 `json:"template_name"`
	Data         map[string]interface{} `json:"data"`
	Channels     []string               `json:"channels"`
	mu           sync.Mutex
}

func NewNotification(recipient Recipient, title, templateName string, channels []string) *Notification {
	if len(channels) == 0 {
		channels = []string{NotificationChannelEmail}
	}
	return &Notification{
		Recipient:    recipient,
		Title:        title,
		TemplateName: templateName,
		Data:         make(map[string]any),
		Channels:     channels,
	}
}

func (n *Notification) AddData(key string, value any) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Data[key] = value
}

func (n *Notification) RemoveData(key string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.Data, key)
}

func (n *Notification) GetChannels() []string {
	return n.Channels
}

func (n *Notification) AddChannel(channel string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Channels = append(n.Channels, channel)
}

func (n *Notification) AddChannels(channels []string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Channels = append(n.Channels, channels...)
}

func (n *Notification) RemoveChannel(channel string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for i, c := range n.Channels {
		if c == channel {
			n.Channels = append(n.Channels[:i], n.Channels[i+1:]...)
		}
	}
}
