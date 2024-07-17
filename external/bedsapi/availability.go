package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type AvailabilityRequestStay struct {
	Checkin        string `json:"checkIn"`
	Checkout       string `json:"checkOut"`
	Shiftdays      int32  `json:"shiftDays,omitempty"`
	AllowOnlyShift bool   `json:"allowOnlyShift,omitempty"`
}

type AvailabilityRequestGeoLocation struct {
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	Radius             string `json:"radius,omitempty"`
	Uint               string `json:"uint,omitempty"`
	SecondaryLatitude  string `json:"secondaryLatitude,omitempty"`
	SecondaryLongitude string `json:"secondaryLongitude,omitempty"`
}

type AvailabilityRequestFilter struct {
	MaxHotels       int32  `json:"maxHotels,omitempty"`
	MaxRooms        int32  `json:"maxRooms,omitempty"`
	MinRate         string `json:"minRate,omitempty"`
	MaxRate         string `json:"maxRate,omitempty"`
	MaxRatesPerRoom int32  `json:"maxRatesPerRoom,omitempty"`
	MinCategory     int32  `json:"minCategory,omitempty"`
	MaxCategory     int32  `json:"maxCategory,omitempty"`
	Contract        string `json:"contract,omitempty"`
}

type AvailabilityRequestBoards struct {
	Board    []string `json:"board"`
	Included bool     `json:"included"`
}

type AvailabilityRequestRooms struct {
	Room     []string `json:"room"`
	Included bool     `json:"included"`
}

type AvailabilityRequestPaxes struct {
	Type string `json:"type"`
	Age  int32  `json:"age"`
}

type AvailabilityRequestOccupancies struct {
	Rooms    int32                      `json:"rooms"`
	Adults   int32                      `json:"adults"`
	Children int32                      `json:"children"`
	Paxes    []AvailabilityRequestPaxes `json:"paxes,omitempty"`
}

type AvailabilityRequestKeywords struct {
	AllIncluded bool    `json:"allIncluded,omitempty"`
	Keyword     []int32 `json:"keyword,omitempty"`
}

type AvailabilityRequestHotels struct {
	Hotel []int32 `json:"hotel,omitempty"`
}

type AvailabilityRequestReview struct {
	Type           string  `json:"type"`
	MinRate        *string `json:"minRate,omitempty"`
	MaxRate        *string `json:"maxRate,omitempty"`
	MinReviewCount *int32  `json:"minReviewCount,omitempty"`
}

// type

type AvailabilityRequest struct {
	Stay          AvailabilityRequestStay          `json:"stay"`
	Geolocation   *AvailabilityRequestGeoLocation  `json:"geolocation,omitempty"`
	Filter        *AvailabilityRequestFilter       `json:"filter,omitempty"`
	Boards        *AvailabilityRequestBoards       `json:"boards,omitempty"`
	Rooms         *AvailabilityRequestRooms        `json:"rooms,omitempty"`
	DailyRate     *bool                            `json:"dailyRate,omitempty"`
	SourceMarket  *string                          `json:"sourceMarket,omitempty"`
	AifUse        *bool                            `json:"aifUse,omitempty"`
	Platform      *int32                           `json:"platform,omitempty"`
	Language      *string                          `json:"language,omitempty"`
	Occupancies   []AvailabilityRequestOccupancies `json:"occupancies"`
	Keywords      *AvailabilityRequestKeywords     `json:"keywords,omitempty"`
	Hotels        *AvailabilityRequestHotels       `json:"hotels,omitempty"`
	Review        *[]AvailabilityRequestReview     `json:"review,omitempty"`
	Accomodations *[]string                        `json:"accomodations,omitempty"`
	Inclusions    *[]string                        `json:"inclusions,omitempty"`
}

type AvailabilityResponse struct {
	AuditData struct {
		ProcessTime *string `json:"processTime,omitempty"`
		Timestamp   *string `json:"timestamp,omitempty"`
		RequestHost *string `json:"requestHost,omitempty"`
		ServerID    *string `json:"serverId,omitempty"`
		Enviroment  *string `json:"enviroment,omitempty"`
		Release     *string `json:"release,omitempty"`
		Token       *string `json:"token,omitempty"`
	} `json:"auditData"`
	Error *struct {
		Code    string  `json:"code"`
		Message *string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	Hotels *struct {
		CheckIn  *string `json:"checkIn,omitempty"`
		CheckOut *string `json:"checkOut,omitempty"`
		Hotels   *[]struct {
			CheckOut         *string `json:"checkOut,omitempty"`
			CheckIn          *string `json:"checkIn,omitempty"`
			Code             *int32  `json:"code,omitempty"`
			Name             *string `json:"name,omitempty"`
			Description      *string `json:"description,omitempty"`
			ExclusiveDeal    *int32  `json:"exclusiveDeal,omitempty"`
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
			TotalSellingRate *int    `json:"totalSellingRate,omitempty"`
			TotalNet         *int    `json:"totalNet,omitempty"`
			PendingAmount    *int    `json:"pendingAmount,omitempty"`
			Currency         *string `json:"currency,omitempty"`
			Supplier         *struct {
				Name      *string `json:"name,omitempty"`
				VatNumber *string `json:"vatNumber,omitempty"`
			} `json:"supplier,omitempty"`
			ClientComments     *string `json:"clientComments,omitempty"`
			CancellationAMount *string `json:"cancellationAMount,omitempty"`
			Upselling          *struct {
				Rooms *[]struct {
					Status           *string `json:"status,omitempty"`
					ID               *int32  `json:"id,omitempty"`
					Code             *string `json:"code,omitempty"`
					Name             *string `json:"name,omitempty"`
					SupplierRefrence *string `json:"supplierRefrence,omitempty"`
					Paxes            *[]struct {
						RoomID  *int32  `json:"roomId,omitempty"`
						Type    string  `json:"type"`
						Age     *int32  `json:"age,omitempty"`
						Name    *string `json:"name,omitempty"`
						Surname *string `json:"surname,omitempty"`
					} `json:"paxes,omitempty"`
					Rates *[]struct {
						RateKey              *string `json:"rateKey,omitempty"`
						RateClass            *string `json:"rateClass,omitempty"`
						RateType             *string `json:"rateType,omitempty"`
						Net                  *string `json:"net,omitempty"`
						Discount             *string `json:"discount,omitempty"`
						DiscountPCT          *string `json:"discountPCT,omitempty"`
						SellingRate          *string `json:"sellingRate,omitempty"`
						HotelMandatory       *bool   `json:"hotelMandatory,omitempty"`
						Allotment            *int32  `json:"allotment,omitempty"`
						Commission           *string `json:"commission,omitempty"`
						CommissionVAT        *string `json:"commissionVAT,omitempty"`
						CommissionPCT        *string `json:"commissionPCT,omitempty"`
						RateCommentsID       *string `json:"rateCommentsId,omitempty"`
						RateComments         *string `json:"rateComments,omitempty"`
						Packaging            *bool   `json:"packaging,omitempty"`
						BoardCode            *string `json:"boardCode,omitempty"`
						BoardName            *string `json:"boardName,omitempty"`
						Rooms                *int32  `json:"rooms,omitempty"`
						Adults               *int32  `json:"adults,omitempty"`
						Children             *int32  `json:"children,omitempty"`
						ChildrenAges         *string `json:"childrenAges,omitempty"`
						Rateup               *string `json:"rateup,omitempty"`
						Resident             *bool   `json:"resident,omitempty"`
						CancellationPolicies *[]struct {
							Amount         *string `json:"amount,omitempty"`
							From           *string `json:"from,omitempty"`
							Percent        *string `json:"percent,omitempty"`
							NumberOfNights *int32  `json:"numberOfNights,omitempty"`
						} `json:"cancellationPolicies,omitempty"`
						Taxes struct {
							AllIncluded *bool   `json:"allIncluded,omitempty"`
							TaxScheme   *string `json:"taxScheme,omitempty"`
							Taxes       *[]struct {
								Included       *bool   `json:"included,omitempty"`
								Percent        *string `json:"percent,omitempty"`
								Amount         *string `json:"amount,omitempty"`
								Currency       *string `json:"currency,omitempty"`
								Type           *string `json:"type,omitempty"`
								SubType        *string `json:"subType,omitempty"`
								Clientamount   *string `json:"clientamount,omitempty"`
								ClientCurrency *string `json:"clientCurrency,omitempty"`
							} `json:"taxes,omitempty"`
						} `json:"taxes,omitempty"`
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
						ShiftRates *[]struct {
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
							CheckIn        *string `json:"checkIn,omitempty"`
							CheckOut       *string `json:"checkOut,omitempty"`
							Brand          *string `json:"brand,omitempty"`
							Resident       *bool   `json:"resident,omitempty"`
						} `json:"shiftRates,omitempty"`
						DailyRates *[]struct {
							Offset           *int32  `json:"offset,omitempty"`
							DailyNet         *string `json:"dailyNet,omitempty"`
							DailySellingRate *string `json:"dailySellingRate,omitempty"`
						} `json:"dailyRates,omitempty"`
					} `json:"rates,omitempty"`
				} `json:"rooms,omitempty"`
			} `json:"upselling,omitempty"`
		} `json:"hotels,omitempty"`
	} `json:"hotels,omitempty"`
}

func (b *hotelbeds) GetHotels(req AvailabilityRequest) (*AvailabilityResponse, error) {
	res, err := b.makeRequest(http.MethodPost, fmt.Sprintf("%s/hotels", b.baseUrl), req, nil)
	if err != nil {
		return nil, err
	}

	var resp AvailabilityResponse
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
