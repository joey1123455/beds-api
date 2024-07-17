package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type ApiRateBookingList struct {
	Amount    *string `json:"amount,omitempty"`
	BoardCode *string `json:"boardCode,omitempty"`
	Rooms     *int32  `json:"rooms,omitempty"`
}

type RoomBookingList struct {
	Status *string               `json:"status,omitempty"`
	Id     *int32                `json:"id,omitempty"`
	Code   *string               `json:"code,omitempty"`
	Paxes  *[]ApiPax             `json:"paxes,omitempty"`
	Rates  *[]ApiRateBookingList `json:"rates,omitempty"`
}

type ApiHotelBookingList struct {
	CheckOut           *string          `json:"checkOut,omitempty"`
	CheckIn            *string          `json:"checkIn,omitempty"`
	Code               *int32           `json:"code,omitempty"`
	Name               *string          `json:"name,omitempty"`
	DestinationCode    *string          `json:"destinationCode,omitempty"`
	CancellationAmount *string          `json:"cancellationAmount"`
	Rooms              *RoomBookingList `json:"rooms,omitempty"`
}

type ApiBookingList struct {
	Reference        *string              `json:"reference,omitempty"`
	ClientReference  *string              `json:"clientReference,omitempty"`
	Status           *string              `json:"status,omitempty"`
	CreationUser     *string              `json:"creationUser,omitempty"`
	Holdere          *ApiHolder           `json:"holder,omitempty"`
	Hotel            *ApiHotelBookingList `json:"hotel,omitempty"`
	InvoiceCompany   *APiReceptive        `json:"invoiceCompany,omitempty"`
	TotalSellingRate *string              `json:"totalSellingRate,omitempty"`
	TotalNet         *string              `json:"totalNet,omitempty"`
	PendingAmount    *int                 `json:"pendingAmount,omitempty"`
	Currency         *string              `json:"currency,omitempty"`
}

type ApiBookingsList struct {
	Bookings *[]ApiBookingList `json:"bookings"`
}

type BookingListResponse struct {
	AuditData APiAuditData `json:"auditData"`
	Error     *ApiError    `json:"error"`
	Bookings  *ApiBookingsList
}

func (b *hotelbeds) ListBookings(queries map[string]string) (*BookingListResponse, error) {
	_, ok := queries["from"]
	if !ok {
		return nil, fmt.Errorf("from is required")
	}
	_, ok = queries["to"]
	if !ok {
		return nil, fmt.Errorf("to is required")
	}

	_, ok = queries["start"]
	if !ok {
		return nil, fmt.Errorf("start is required")
	}
	_, ok = queries["end"]
	if !ok {
		return nil, fmt.Errorf("end is required")
	}

	res, err := b.makeRequest(http.MethodGet, fmt.Sprintf("%s/bookings", b.baseUrl), map[string]any{}, queries)
	if err != nil {
		return nil, err
	}

	var resp BookingListResponse
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
