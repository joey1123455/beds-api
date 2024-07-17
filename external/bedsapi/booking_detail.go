package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type ApiHotel struct {
	CheckOut            *string          `json:"checkOut,omitempty"`
	CheckIn             *string          `json:"checkIn,omitempty"`
	Code                *int32           `json:"code,omitempty"`
	Name                *string          `json:"name,omitempty"`
	Description         *string          `json:"description,omitempty"`
	ExclusiveDeal       *int32           `json:"exclusiveDeal,omitempty"`
	CategoryCode        *string          `json:"categoryCode,omitempty"`
	CategoryName        *string          `json:"categoryName,omitempty"`
	DestinationCode     *string          `json:"destinationCode,omitempty"`
	DestinationName     *string          `json:"destinationName,omitempty"`
	ZoneCode            *int32           `json:"zoneCode,omitempty"`
	ZoneName            *string          `json:"zoneName,omitempty"`
	Latitude            *string          `json:"latitude,omitempty"`
	Longitude           *string          `json:"longitude,omitempty"`
	MinRate             *string          `json:"minRate,omitempty"`
	MaxRate             *string          `json:"maxRate,omitempty"`
	TotalSellingRate    *int             `json:"totalSellingRate,omitempty"`
	PendingAmount       *int             `json:"pendingAmount,omitempty"`
	TotalNet            *int             `json:"totalNet,omitempty"`
	Currency            *string          `json:"currency,omitempty"`
	Supplier            *ApiSupplier     `json:"supplier,omitempty"`
	ClientComments      *string          `json:"clientComments,omitempty"`
	CancellationAmount  *string          `json:"cancellationAmount,omitempty"`
	Upselling           *APiUpselling    `json:"upselling,omitempty"`
	Keywords            *[]ApiKeyword    `json:"keyWords,omitempty"`
	Reviews             *[]ApiReviews    `json:"reviews,omitempty"`
	Rooms               *[]Room          `json:"rooms,omitempty"`
	CreditCards         *[]ApiCreditCard `json:"creditCards,omitempty"`
	PaymentdataRequired *bool            `json:"paymentdataRequired,omitempty"`
}

type ApiBooking struct {
	Reference             *string                  `json:"reference,omitempty"`
	CancellationReference *string                  `json:"cancellationReference,omitempty"`
	CreationDate          *string                  `json:"creationDate,omitempty"`
	Status                *string                  `json:"status,omitempty"`
	ModificationPolicies  *ApiModificationPolicies `json:"modificationPolicies,omitempty"`
	AgCommision           *string                  `json:"agCommision,omitempty"`
	CommisionVAT          *string                  `json:"commisionVAT,omitempty"`
	CreationUser          *string                  `json:"creationUser,omitempty"`
	Holder                *ApiHolder               `json:"holder,omitempty"`
	Remark                *string                  `json:"remark,omitempty"`
	InvoiceCompany        *APiReceptive            `json:"invoiceCompany,omitempty"`
	TotalSellingRate      *int                     `json:"totalSellingRate,omitempty"`
	TotalNet              *int                     `json:"totalNet,omitempty"`
	PendingAmount         *int                     `json:"pendingAmount,omitempty"`
	Currency              *string                  `json:"currency,omitempty"`
	Hotel                 *ApiHotel                `json:"hotel,omitempty"`
}

type BookingDetailResponse struct {
	AuditData APiAuditData `json:"auditData"`
	Error     *ApiError    `json:"error"`
	Booking   *ApiBooking  `json:"booking"`
}

func (b *hotelbeds) BookingDetail(reference string) (*BookingDetailResponse, error) {
	res, err := b.makeRequest(http.MethodGet, fmt.Sprintf("%s/bookings/%s", b.baseUrl, reference), map[string]any{}, nil)
	if err != nil {
		return nil, err
	}

	var resp BookingDetailResponse
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
