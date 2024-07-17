package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type BookingReconfirmationResponse struct {
	AuditData APiAuditData  `json:"auditData"`
	Error     *ApiError     `json:"error"`
	Booking   *[]ApiBooking `json:"booking"`
}

func (b *hotelbeds) ReconfirmBooking(queries map[string]string) (*BookingReconfirmationResponse, error) {
	_, ok := queries["from"]
	if !ok {
		return nil, fmt.Errorf("from is required")
	}
	_, ok = queries["to"]
	if !ok {
		return nil, fmt.Errorf("to is required")
	}

	res, err := b.makeRequest(http.MethodGet, fmt.Sprintf("%s/bookings/reconfirmations", b.baseUrl), map[string]any{}, queries)
	if err != nil {
		return nil, err
	}

	var resp BookingReconfirmationResponse
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
