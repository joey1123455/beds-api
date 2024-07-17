package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type BookingCancellationResponse struct {
	AuditData APiAuditData `json:"auditData"`
	Error     *ApiError    `json:"error"`
	Booking   *ApiBooking  `json:"booking"`
}

func (b *hotelbeds) CancelBooking(ref string, queries map[string]string) (*BookingCancellationResponse, error) {
	res, err := b.makeRequest(http.MethodDelete, fmt.Sprintf("%s/bookings/%s", b.baseUrl, ref), map[string]any{}, queries)
	if err != nil {
		return nil, err
	}

	var resp BookingCancellationResponse
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
