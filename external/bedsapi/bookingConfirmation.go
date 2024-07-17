package bedsapi

import (
	"fmt"
	"net/http"

	"github.com/joey1123455/beds-api/internal/httplib"
)

type ApiHolder struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type BookingConfirmationRequestPaymentCard struct {
	CardType       string `json:"cardType"`
	CardNumber     string `json:"cardNumber"`
	CardHolderName string `json:"cardHolderName"`
	ExpiryDate     string `json:"expiryDate"`
	CardCVC        string `json:"cardCVC"`
}

type BookingConfirmationRequestContactData struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type BookingConfirmationRequestBillingAddress struct {
	Address1    string  `json:"address1"`
	Address2    *string `json:"address2,omitempty"`
	City        string  `json:"city"`
	State       *string `json:"state,omitempty"`
	PostalCode  string  `json:"postalCode"`
	CountryCode string  `json:"countryCode"`
}

type BookingConfirmationRequestThreeDsDataInfoProvided struct {
	Id   string `json:"id"`
	Cavv string `json:"cavv"`
	Eci  string `json:"eci"`
}

type BookingConfirmationRequestThreeDsData struct {
	Option       string                                             `json:"option"`
	Version      string                                             `json:"version"`
	InfoProvided *BookingConfirmationRequestThreeDsDataInfoProvided `json:"infoProvided,omitempty"`
}

type BookingConfirmationRequestDevice struct {
	Id        *string `json:"id,omitempty"`
	Ip        string  `json:"ip"`
	UserAgent string  `json:"userAgent"`
}

type BookingConfirmationRequestPaymentData struct {
	PaymentCard    BookingConfirmationRequestPaymentCard     `json:"paymentCard"`
	ContactData    BookingConfirmationRequestContactData     `json:"contactData"`
	BillingAddress *BookingConfirmationRequestBillingAddress `json:"billingAddress,omitempty"`
	ThreeDsData    *BookingConfirmationRequestThreeDsData    `json:"threeDsData,omitempty"`
	WebPatner      *int32                                    `json:"webPatner,omitempty"`
	Device         *BookingConfirmationRequestDevice         `json:"device,omitempty"`
}

type BookingConfirmationRequestEmail struct {
	To    *string `json:"to,omitempty"`
	From  *string `json:"from,omitempty"`
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}

type BookingConfirmationRequestVoucher struct {
	Language *string                          `json:"language,omitempty"`
	Email    *BookingConfirmationRequestEmail `json:"email,omitempty"`
	Logo     *string                          `json:"logo,omitempty"`
}

type BookingConfirmationRequestRoomPaxes struct {
	RoomID  *int32  `json:"roomId,omitempty"`
	Age     *int32  `json:"age,omitempty"`
	Name    *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
}

type BookingConfirmationRequestRoom struct {
	RateKey string    `json:"rateKey"`
	Paxes   *[]ApiPax `json:"paxes,omitempty"`
}

type BookingConfirmationRequest struct {
	Holder          ApiHolder                              `json:"holder"`
	PaymentData     *BookingConfirmationRequestPaymentData `json:"paymentData,omitempty"`
	ClientReference string                                 `json:"clientReference"`
	CreationUser    *string                                `json:"creationUser,omitempty"`
	Remark          *string                                `json:"remark,omitempty"`
	Voucher         *BookingConfirmationRequestVoucher     `json:"voucher,omitempty"`
	Tolerance       *string                                `json:"tolerance,omitempty"`
	Language        *string                                `json:"language,omitempty"`
	Rooms           *[]BookingConfirmationRequestRoom      `json:"rooms,omitempty"`
}

type APiAuditData struct {
	ProcessTime *string `json:"processTime,omitempty"`
	Timestamp   *string `json:"timestamp,omitempty"`
	RequestHost *string `json:"requestHost,omitempty"`
	ServerID    *string `json:"serverId,omitempty"`
	Enviroment  *string `json:"enviroment,omitempty"`
	Release     *string `json:"release,omitempty"`
	Token       *string `json:"token,omitempty"`
}

type ApiError struct {
	Code    string  `json:"code"`
	Message *string `json:"message,omitempty"`
}

type ApiModificationPolicies struct {
	Modification *string `json:"modification,omitempty"`
	Cancellation *string `json:"cancellation,omitempty"`
}

type APiReceptive struct {
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	Code               *string `json:"code,omitempty"`
	Name               *string `json:"name,omitempty"`
}

type ApiSupplier struct {
	Name      *string `json:"name,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
}

type ApiPax struct {
	RoomId  *int32  `json:"roomId,omitempty"`
	Type    string  `json:"type"`
	Age     *int32  `json:"age,omitempty"`
	Name    *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
}

type CancellationPolicy struct {
	Amount         *string `json:"amount,omitempty"`
	From           *string `json:"from,omitempty"`
	Percent        *string `json:"percent,omitempty"`
	NumberOfNights *string `json:"numberOfNights,omitempty"`
}

type ApiTax struct {
	Included       *bool   `json:"included,omitempty"`
	Percent        *string `json:"percent,omitempty"`
	Amount         *string `json:"amount,omitempty"`
	Currency       *string `json:"currency,omitempty"`
	Type           *string `json:"type,omitempty"`
	SubType        *string `json:"subType,omitempty"`
	ClientAmount   *string `json:"clientAmount,omitempty"`
	ClientCurrency *string `json:"clientCurrency,omitempty"`
}

type ApiTaxes struct {
	AllIncluded *bool     `json:"allIncluded,omitempty"`
	TaxScheme   *string   `json:"taxScheme,omitempty"`
	Taxes       *[]ApiTax `json:"taxes,omitempty"`
}

type ApiPromotion struct {
	Code   *string `json:"code,omitempty"`
	Name   *string `json:"name,omitempty"`
	Remark *string `json:"remark,omitempty"`
}

type ApiOffer struct {
	Code   *string `json:"code,omitempty"`
	Name   *string `json:"name,omitempty"`
	AMount *string `json:"amount,omitempty"`
}

type ApiShiftRate struct {
	RateKey        *string `json:"rateKey,omitempty"`
	RateClass      *string `json:"rateClass,omitempty"`
	RateType       *string `json:"rateType,omitempty"`
	Net            *string `json:"net,omitempty"`
	Discount       *string `json:"discount,omitempty"`
	DiscountPCT    *string `json:"discountPCT,omitempty"`
	SellingRate    *string `json:"sellingRate,omitempty"`
	HotelMandatory *bool   `json:"hotelMandatory,omitempty"`
	Allotment      *int32  `json:"allotment,omitempty"`
	CommissionVAT  *string `json:"commissionVAT,omitempty"`
	CommissionPCT  *string `json:"commissionPCT,omitempty"`
	CheckIn        *string `json:"checkIn,omitempty"`
	CheckOut       *string `json:"checkOut,omitempty"`
	Brand          *string `json:"brand,omitempty"`
	Resident       *bool   `json:"resident,omitempty"`
}

type ApiDailyRate struct {
	Offset           *int32  `json:"offset,omitempty"`
	DailyNet         *string `json:"dailyNet,omitempty"`
	DailySellingRate *string `json:"dailySellingRate,omitempty"`
}

type ApiRate struct {
	RateKey              *string               `json:"rateKey,omitempty"`
	RateClass            *string               `json:"rateClass,omitempty"`
	RateType             *string               `json:"rateType,omitempty"`
	Net                  *string               `json:"net,omitempty"`
	Discount             *string               `json:"discount,omitempty"`
	DiscountPCT          *string               `json:"discountPCT,omitempty"`
	SellingRate          *string               `json:"sellingRate,omitempty"`
	HotelMandatory       *bool                 `json:"hotelMandatory,omitempty"`
	Allotment            *int32                `json:"allotment,omitempty"`
	Commission           *string               `json:"commission,omitempty"`
	CommissionVAT        *string               `json:"commissionVAT,omitempty"`
	CommissionPCT        *string               `json:"commissionPCT,omitempty"`
	RateCommentsId       *string               `json:"rateCommentsId,omitempty"`
	RateComments         *string               `json:"rateComments,omitempty"`
	Packaging            *bool                 `json:"packaging,omitempty"`
	BoardCode            *string               `json:"boardCode,omitempty"`
	Rooms                *int32                `json:"rooms,omitempty"`
	Adults               *int32                `json:"adults,omitempty"`
	Children             *int32                `json:"children,omitempty"`
	ChildrenAges         *string               `json:"childrenAges,omitempty"`
	RateUp               *string               `json:"rateUp,omitempty"`
	Brand                *string               `json:"brand,omitempty"`
	Resident             *bool                 `json:"resident,omitempty"`
	CancellationPolicies *[]CancellationPolicy `json:"cancellationPolicies,omitempty"`
	Taxes                *[]ApiTaxes           `json:"taxes,omitempty"`
	Promotions           *[]ApiPromotion       `json:"promotions,omitempty"`
	Offers               *[]ApiOffer           `json:"offers,omitempty"`
	ShiftRates           *[]ApiShiftRate       `json:"shiftRates,omitempty"`
	DailyRates           *[]ApiDailyRate       `json:"dailyRates,omitempty"`
}

type Room struct {
	Status            *string    `json:"status,omitempty"`
	Id                *int32     `json:"id,omitempty"`
	Code              *string    `json:"code,omitempty"`
	Name              *string    `json:"name,omitempty"`
	SupplierReference *string    `json:"supplierReference,omitempty"`
	Paxes             *[]ApiPax  `json:"paxes,omitempty"`
	Rates             *[]ApiRate `json:"rates,omitempty"`
}

type APiUpselling struct {
	Rooms *[]Room `json:"rooms,omitempty"`
}

type ApiKeyword struct {
	Code   string `json:"code,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
}

type ApiReviews struct {
	Rate        *string `json:"rate,omitempty"`
	ReviewCount *int32  `json:"reviewCount,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type ApiCancellationPolicies struct {
	Amount         *string `json:"amount,omitempty"`
	From           *string `json:"from,omitempty"`
	To             *string `json:"to,omitempty"`
	Percentage     *string `json:"Percentage,omitempty"`
	NumberOfNights *string `json:"numberOfNights,omitempty"`
}

type ApiRateBookingConfirm struct {
	RateKey              *string                    `json:"rateKey,omitempty"`
	RateClass            *string                    `json:"rateClass,omitempty"`
	RateType             *string                    `json:"rateType,omitempty"`
	Net                  *string                    `json:"net,omitempty"`
	Discount             *string                    `json:"discount,omitempty"`
	DiscountPCT          *string                    `json:"discountPCT,omitempty"`
	SellingRate          *string                    `json:"sellingRate,omitempty"`
	HotelMandatory       *bool                      `json:"hotelMandatory,omitempty"`
	Allotment            *int32                     `json:"allotment,omitempty"`
	Commission           *string                    `json:"commission,omitempty"`
	CommissionVat        *string                    `json:"commissionVat,omitempty"`
	CommissionPCT        *string                    `json:"commissionPCT,omitempty"`
	RateCommentsId       *string                    `json:"rateCommentsId,omitempty"`
	RateComments         *string                    `json:"rateComments,omitempty"`
	Packaging            *string                    `json:"packaging,omitempty"`
	BoardCode            *string                    `json:"boardCode,omitempty"`
	BoardName            *string                    `json:"boardName,omitempty"`
	Rooms                *int32                     `json:"rooms,omitempty"`
	Adults               *int32                     `json:"adults,omitempty"`
	Children             *int32                     `json:"children,omitempty"`
	ChildrenAges         *string                    `json:"childrenAges,omitempty"`
	Rateup               *string                    `json:"rateup,omitempty"`
	Brand                *string                    `json:"brand,omitempty"`
	Resident             *bool                      `json:"resident,omitempty"`
	CancellationPolicies *[]ApiCancellationPolicies `json:"cancellationPolicies,omitempty"`
	Taxes                *ApiTaxes                  `json:"taxes,omitempty"`
	Promotions           *[]ApiPromotion            `json:"promotions,omitempty"`
	Offers               *[]ApiOffer                `json:"offers,omitempty"`
	ShiftRates           *[]ApiShiftRate            `json:"shiftRates,omitempty"`
}

type RoomBookingConfirmation struct {
	Status            *string                  `json:"status,omitempty"`
	Id                *string                  `json:"id,omitempty"`
	Code              *string                  `json:"code,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	SupplierReference *string                  `json:"supplierReference,omitempty"`
	Paxes             *[]ApiPax                `json:"paxes,omitempty"`
	Rates             *[]ApiRateBookingConfirm `json:"rates,omitempty"`
}

type ApiCreditCard struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

type APiHotelBookingConfirm struct {
	CheckOut            *string                    `json:"checkOut,omitempty"`
	CheckIn             *string                    `json:"checkIn,omitempty"`
	Code                *int32                     `json:"code,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	CategoryCode        *string                    `json:"categoryCode,omitempty"`
	CategoryName        *string                    `json:"categoryName,omitempty"`
	DestinationCode     *string                    `json:"destinationCode,omitempty"`
	DestinationName     *string                    `json:"destinationName,omitempty"`
	ZoneCode            *int32                     `json:"zoneCode,omitempty"`
	ZoneName            *string                    `json:"zoneName,omitempty"`
	Latitude            *string                    `json:"latitude,omitempty"`
	Longitude           *string                    `json:"longitude,omitempty"`
	MinRate             *string                    `json:"minRate,omitempty"`
	MaxRate             *string                    `json:"maxRate,omitempty"`
	TotalSellingRate    *string                    `json:"totalSellingRate,omitempty"`
	TotalNet            *string                    `json:"totalNet,omitempty"`
	PendingAmount       *int                       `json:"pendingAmount,omitempty"`
	Currency            *string                    `json:"currency,omitempty"`
	Supplier            *ApiSupplier               `json:"supplier,omitempty"`
	ClientComments      *string                    `json:"clientComments,omitempty"`
	CancellationAmount  *string                    `json:"cancellationAmount,omitempty"`
	Upselling           *APiUpselling              `json:"upselling,omitempty"`
	KeyWords            *[]ApiKeyword              `json:"keyWords,omitempty"`
	Reviews             *[]ApiReviews              `json:"reviews,omitempty"`
	Rooms               *[]RoomBookingConfirmation `json:"rooms,omitempty"`
	CreditCards         *[]ApiCreditCard           `json:"creditCards,omitempty"`
	PaymentdataRequired *bool                      `json:"paymentdataRequired,omitempty"`
}

type ApiBookingData struct {
	Reference             *string                  `json:"reference,omitempty"`
	CancellationReference *string                  `json:"cancellationReference,omitempty"`
	ClientReference       *string                  `json:"clientReference,omitempty"`
	CreationDate          *string                  `json:"creationDate,omitempty"`
	Status                *string                  `json:"status,omitempty"`
	ModificationPolicies  *ApiModificationPolicies `json:"modificationPolicies,omitempty"`
	AgCommision           *string                  `json:"agCommision,omitempty"`
	CommisionVAT          *string                  `json:"commisionVAT,omitempty"`
	CreationUser          *string                  `json:"creationUser,omitempty"`
	Holder                *ApiBookingData          `json:"holder,omitempty"`
	Remark                *string                  `json:"remark,omitempty"`
	InvoiceCompany        *APiReceptive            `json:"invoiceCompany,omitempty"`
	TotalSellingRate      *int                     `json:"totalSellingRate,omitempty"`
	TotalNet              *int                     `json:"totalNet,omitempty"`
	PendingAmount         *int                     `json:"pendingAmount,omitempty"`
	Currency              *string                  `json:"currency,omitempty"`
	Hotel                 *APiHotelBookingConfirm  `json:"hotel,omitempty"`
}

type BookingConfirmationResponse struct {
	AuditData APiAuditData    `json:"auditData"`
	Error     *ApiError       `json:"error,omitempty"`
	Booking   *ApiBookingData `json:"booking,omitempty"`
}

func (b *hotelbeds) ConfirmBooking(req BookingConfirmationRequest) (*BookingConfirmationResponse, error) {
	res, err := b.makeRequest(http.MethodPost, fmt.Sprintf("%s/bookings", b.baseUrl), req, nil)
	if err != nil {
		return nil, err
	}

	var resp BookingConfirmationResponse
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
