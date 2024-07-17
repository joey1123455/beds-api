package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type ApiPaymentCard struct {
	CardType       string `json:"cardType"`
	CardNumber     string `json:"cardNumber"`
	CardHolderName string `json:"cardHolderName"`
	ExpiryDate     string `json:"expiryDate"`
	CardCVC        string `json:"cardCVC"`
}

type ApiPaymentContactData struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type ApiBillingAddress struct {
	Address1    string  `json:"address1"`
	Address2    *string `json:"address2,omitempty"`
	City        string  `json:"city"`
	State       *string `json:"state,omitempty"`
	PostalCode  string  `json:"postalCode"`
	CountryCode string  `json:"countryCode"`
}

type InfoProvided struct {
	Id   string `json:"id"`
	Cavv string `json:"cavv"`
	Eci  string `json:"eci"`
}

type ThreeDsData struct {
	Option       string        `json:"option"`
	Version      string        `json:"version"`
	InfoProvided *InfoProvided `json:"infoProvided,omitempty"`
}

type ApiBookingDevice struct {
	Id        *string `json:"id,omitempty"`
	Ip        string  `json:"ip"`
	UserAgent string  `json:"userAgent"`
}

type ApiPaymentData struct {
	PaymentCard    ApiPaymentCard        `json:"paymentCard"`
	ContactData    ApiPaymentContactData `json:"contactData"`
	BillingAddress *ApiBillingAddress    `json:"billingAddress,omitempty"`
	ThreeDsData    *ThreeDsData          `json:"threeDsData,omitempty"`
	WebPatner      *int32                `json:"webPatner,omitempty"`
	Device         *ApiBookingDevice     `json:"device,omitempty"`
}

type BookingChangeRequest struct {
	Mode        string          `json:"mode"`
	PaymentData *ApiPaymentData `json:"paymentData,omitempty"`
	Language    *string         `json:"language,omitempty"`
	Booking     ApiBooking      `json:"booking"`
}

type BookingChangeResponse struct {
	AuditData APiAuditData `json:"auditData"`
	Error     *ApiError    `json:"error"`
	Booking   *ApiBooking  `json:"booking"`
}

func (b *hotelbeds) ChangeBooking(ref string, req BookingChangeRequest) (*BookingChangeResponse, error) {
	res, err := b.makeRequest(http.MethodPost, fmt.Sprintf("%s/bookings/%s", b.baseUrl, ref), req, nil)
	if err != nil {
		return nil, err
	}

	var resp BookingChangeResponse
	var errResp string
	switch res.StatusCode {
	case http.StatusOK, http.StatusInternalServerError, http.StatusBadRequest:
		if err := httplib.ReadJsonResponse(res, &resp); err != nil {
			return nil, err
		}

		return &resp, nil

	case http.StatusUnauthorized, http.StatusForbidden:
		if err := httplib.ReadJsonResponse(res, &errResp); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errResp)

	default:
		return nil, fmt.Errorf("no response body received: %v", res.StatusCode)
	}
}
