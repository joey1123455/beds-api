package bedsapi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// )

func initBeds() (BedsSdk, error) {

	err := godotenv.Load("../../.env.test")
	if err != nil {
		return nil, err
	}

	sdk := New(os.Getenv("HOTEL_BEDS_SECRET_KEY"), os.Getenv("HOTEL_BEDS_API_KEY"), os.Getenv("HOTEL_BEDS_BASE_URL"), nil)
	return sdk, err
}

func Test_Availability(t *testing.T) {
	sdk, err := initBeds()
	assert.Nil(t, err)
	req := AvailabilityRequest{
		Stay: AvailabilityRequestStay{
			Checkin:  "2024-09-15",
			Checkout: "2024-10-15",
		},
		Occupancies: []AvailabilityRequestOccupancies{
			{
				Rooms:    1,
				Adults:   1,
				Children: 0,
			},
		},
		Hotels: &AvailabilityRequestHotels{
			Hotel: []int32{
				123223,
				123224,
				122197,
			},
		},
		// Geolocation: n,
	}

	res, err := sdk.GetHotels(req)
	t.Log(err)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Nil(t, res.Error)
	t.Log(res.Hotels)
}

func Test_BookingConfirmation(t *testing.T) {
	sdk, err := initBeds()
	assert.Nil(t, err)

	roomId := int32(1)
	paxType := "AD"
	firstName := "First Adult Name"
	secondName := "Second Adult Name"
	firstLstName := "Surname"
	secondLstName := "Surname"
	remark := "Booking remarks are to be written here."
	tolerance := "10"

	req := BookingConfirmationRequest{

		Holder: ApiHolder{
			Name:    "HolderFirstName",
			Surname: "HolderLastName",
		},
		Rooms: &[]BookingConfirmationRequestRoom{
			{
				RateKey: "20190315|20190316|W|1|311|DBT.ST|PVP-SHORTSTAY|AI||1~2~0||N@08870BAE87754721542353710729AAES00000010000000007221346",
				Paxes: &[]ApiPax{
					{
						RoomId:  &roomId,
						Type:    paxType,
						Name:    &firstName,
						Surname: &firstLstName,
					},
					{
						RoomId:  &roomId,
						Type:    paxType,
						Name:    &secondName,
						Surname: &secondLstName,
					},
				},
			},
		},
		ClientReference: "IntegrationAgency",
		Remark:          &remark,
		Tolerance:       &tolerance,
	}

	res, err := sdk.ConfirmBooking(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_ListBookings(t *testing.T) {
	sdk, err := initBeds()
	assert.Nil(t, err)

	queries := map[string]string{
		"from":  "1",
		"to":    "30",
		"start": "2024-09-15",
		"end":   "2024-10-15",
	}

	res, err := sdk.ListBookings(queries)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	t.Log(*res.Error.Message)
	assert.Nil(t, res.Error)
}

func Test_BookingDetail(t *testing.T) {
	sdk, err := initBeds()
	assert.Nil(t, err)
	res, err := sdk.BookingDetail("")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	t.Log(*res.Error.Message)
	assert.Nil(t, res.Error)
}

// this endpoint returns 410 error
// func Test_CheckRate(t *testing.T) {
// 	sdk, err := initBeds()
// 	assert.Nil(t, err)
// 	req := CheckRateRequest{
// 		Rooms: &[]CheckRateRequestRoom{
// 			{
// 				RateKey: "20200615|20200616|W|1|311|DBT.ST|PVP-SHORTSTAY|AI||1~2~0||N@08870BAE87754721542353710729AAES00000010000000007221346",
// 			},
// 		},
// 	}

// 	res, err := sdk.CheckRate(req)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, res)
// 	assert.Nil(t, res.Error)
// 	t.Log(res.Error.Message)
// }
