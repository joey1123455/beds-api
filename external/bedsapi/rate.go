package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type CheckRateRequestRoomPaxes struct {
	RoomId  *int32  `json:"roomId,omitempty"`
	Type    string  `json:"type"`
	Age     *int32  `json:"age,omitempty"`
	Name    *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
}

type CheckRateRequestRoom struct {
	RateKey string                     `json:"rateKey"`
	Paxes   *CheckRateRequestRoomPaxes `json:"paxes,omitempty"`
}
type CheckRateRequest struct {
	Upselling *bool                   `json:"upselling,omitempty"`
	ExpandCXL *bool                   `json:"expandCxl,omitempty"`
	Language  *string                 `json:"language,omitempty"`
	Rooms     *[]CheckRateRequestRoom `json:"rooms,omitempty"`
}

type CheckRateResponse struct {
	AuditData struct {
		Token string `json:"token"`
	} `json:"auditData"`
	Error *struct {
		Code    string  `json:"code"`
		Message *string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	Hotels *struct {
		CheckOut         *string `json:"checkOut,omitempty"`
		CheckIn          *string `json:"checkIn,omitempty"`
		Code             *int32  `json:"code,omitempty"`
		Name             *string `json:"name,omitempty"`
		Description      *string `json:"description,omitempty"`
		CategoryCode     *string `json:"categoryCode,omitempty"`
		CategoryName     *string `json:"categoryName,omitempty"`
		DestinationCode  *string `json:"destinationCode,omitempty"`
		DestinationName  *string `json:"destinationName,omitempty"`
		ZoneCode         *int32  `json:"zoneCode,omitempty"`
		ZoneName         *string `json:"zoneName,omitempty"`
		Latitude         *string `json:"latitude,omitempty"`
		Longitude        *string `json:"longitude,omitempty"`
		MinRate          *string `json:"minRate,omitempty"`
		MaxRate          *string `json:"maxRate,omitempty"`
		TotalSellingRate *string `json:"totalSellingRate,omitempty"`
		TotalNet         *string `json:"totalNet,omitempty"`
		Giata            *string `json:"giata,omitempty"`
		Currency         *string `json:"currency,omitempty"`
		Upselling        *struct {
			Rooms *[]struct {
				Rates *[]struct {
					RateKey        *string `json:"rateKey,omitempty"`
					RateClass      *string `json:"rateClass,omitempty"`
					RateType       *string `json:"rateType,omitempty"`
					Net            *string `json:"net,omitempty"`
					Discount       *string `json:"discount,omitempty"`
					DiscountPCT    *string `json:"discountPCT,omitempty"`
					SellingRate    *string `json:"sellingRate,omitempty"`
					HotelMandatory *bool   `json:"hotelMandatory,omitempty"`
					Allotment      *int32  `json:"allotment,omitempty"`
					Commission     *string `json:"commission,omitempty"`
					CommissionVAT  *string `json:"commissionVAT,omitempty"`
					CommissionPCT  *string `json:"commissionPCT,omitempty"`
					RateComments   *string `json:"rateComments,omitempty"`
					Packaging      *bool   `json:"packaging,omitempty"`
					BoardCode      *string `json:"boardCode,omitempty"`
					BoardName      *string `json:"boardName,omitempty"`
					RateBreakDown  *struct {
						RateDiscounts *[]struct {
							Code   *string `json:"code,omitempty"`
							Name   *string `json:"name,omitempty"`
							Amount *string `json:"amount,omitempty"`
						} `json:"rateDiscounts,omitempty"`
						RateSupplements *[]struct {
							Code      *string `json:"code,omitempty"`
							Name      *string `json:"name,omitempty"`
							Amount    *string `json:"amount,omitempty"`
							From      *string `json:"from,omitempty"`
							To        *string `json:"to,omitempty"`
							Nights    *int32  `json:"nights,omitempty"`
							PaxNumber *int32  `json:"paxNumber,omitempty"`
						} `json:"rateSupplements,omitempty"`
					} `json:"rateBreakDown,omitempty"`
					Rooms                *int32  `json:"rooms,omitempty"`
					Adults               *int32  `json:"adults,omitempty"`
					Children             *int32  `json:"children,omitempty"`
					Rateup               *string `json:"rateup,omitempty"`
					Resident             *bool   `json:"resident,omitempty"`
					CancellationPolicies *[]struct {
						Amount *string `json:"amount,omitempty"`
						From   *string `json:"from,omitempty"`
					} `json:"cancellationPolicies,omitempty"`
				} `json:"rates,omitempty"`
			} `json:"rooms,omitempty"`
		} `json:"upselling,omitempty"`
		Rooms *[]struct {
			Code  *string `json:"code,omitempty"`
			Name  *string `json:"name,omitempty"`
			Rates *[]struct {
				RateKey        *string `json:"rateKey,omitempty"`
				RateClass      *string `json:"rateClass,omitempty"`
				RateType       *string `json:"rateType,omitempty"`
				Net            *string `json:"net,omitempty"`
				Discount       *string `json:"discount,omitempty"`
				DiscountPCT    *string `json:"discountPCT,omitempty"`
				SellingRate    *string `json:"sellingRate,omitempty"`
				HotelMandatory *bool   `json:"hotelMandatory,omitempty"`
				Allotment      *int32  `json:"allotment,omitempty"`
				Commission     *string `json:"commission,omitempty"`
				CommissionVAT  *string `json:"commissionVAT,omitempty"`
				CommissionPCT  *string `json:"commissionPCT,omitempty"`
				RateComments   *string `json:"rateComments,omitempty"`
				Packaging      *bool   `json:"packaging,omitempty"`
				BoardCode      *string `json:"boardCode,omitempty"`
				BoardName      *string `json:"boardName,omitempty"`
				RateBreakDown  *struct {
					RateDiscounts *[]struct {
						Code   *string `json:"code,omitempty"`
						Name   *string `json:"name,omitempty"`
						Amount *string `json:"amount,omitempty"`
					}
					RateSupplements *[]struct {
						Code      *string `json:"code,omitempty"`
						Name      *string `json:"name,omitempty"`
						Amount    *string `json:"amount,omitempty"`
						From      *string `json:"from,omitempty"`
						To        *string `json:"to,omitempty"`
						Nights    *int32  `json:"nights,omitempty"`
						PaxNumber *int32  `json:"paxNumber,omitempty"`
					}
				} `json:"rateBreakDown,omitempty"`
				Rooms                *int32 `json:"rooms,omitempty"`
				Adults               *int32 `json:"adults,omitempty"`
				Children             *int32 `json:"children,omitempty"`
				Resident             *bool  `json:"resident,omitempty"`
				CancellationPolicies *[]struct {
					Amount *string `json:"amount,omitempty"`
					From   *string `json:"from,omitempty"`
				} `json:"cancellationPolicies,omitempty"`
				Promotions *[]struct {
					Code   *string `json:"code,omitempty"`
					Name   *string `json:"name,omitempty"`
					Remark *string `json:"remark,omitempty"`
				} `json:"promotions,omitempty"`
				Offers *[]struct {
					Code   *string `json:"code,omitempty"`
					Name   *string `json:"name,omitempty"`
					Amount *string `json:"amount,omitempty"`
				} `json:"offers,omitempty"`
			} `json:"rates,omitempty"`
		} `json:"rooms,omitempty"`
		CreditCards *[]struct {
			Code *string `json:"code,omitempty"`
			Name *string `json:"name,omitempty"`
		} `json:"creditCards,omitempty"`
		PaymentDataRequired  *bool `json:"paymentDataRequired,omitempty"`
		ModificationPolicies *struct {
			Cancellation bool `json:"cancellation,omitempty"`
			Modification bool `json:"modification,omitempty"`
		} `json:"modificationPolicies,omitempty"`
	} `json:"hotels,omitempty"`
}

func (b *hotelbeds) CheckRate(req CheckRateRequest) (*CheckRateResponse, error) {
	res, err := b.makeRequest(http.MethodPost, fmt.Sprintf("%s/checkrates", b.baseUrl), req, nil)
	if err != nil {
		return nil, err
	}

	var resp CheckRateResponse
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
