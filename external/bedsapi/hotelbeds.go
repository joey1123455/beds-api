package bedsapi

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joey1123455/beds-api/internal/httplib"
	"github.com/joey1123455/beds-api/internal/logger"
)

var (
	maxRetry      = 5
	timeout       = 30 * time.Second
	retryWaitTime = 100 * time.Millisecond
)

type BedsSdk interface {
	GetHotels(req AvailabilityRequest) (*AvailabilityResponse, error)
	CheckRate(req CheckRateRequest) (*CheckRateResponse, error)
	ConfirmBooking(req BookingConfirmationRequest) (*BookingConfirmationResponse, error)
	ListBookings(queries map[string]string) (*BookingListResponse, error)
	BookingDetail(reference string) (*BookingDetailResponse, error)
}

type hotelbeds struct {
	SecretKey string
	publickey string
	logger    logger.Logger
	baseUrl   string
}

func New(secretKey, publickey, baseUrl string, logger logger.Logger) BedsSdk {
	return &hotelbeds{SecretKey: secretKey, publickey: publickey, logger: logger, baseUrl: baseUrl}
}

func (k *hotelbeds) makeRequest(method, endpoint string, body interface{}, queries map[string]string) (*http.Response, error) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	signatureString := k.publickey + k.SecretKey + timestamp
	hash := sha256.Sum256([]byte(signatureString))
	signature := fmt.Sprintf("%x", hash)
	log.Println("signature: ", signature)
	client := httplib.NewHttpClient(timeout).
		SetHeaders(map[string]string{
			"Api-key":         k.publickey,
			"accept":          "application/json",
			"content-type":    "application/json",
			"Accept-Encoding": "gzip",
			"X-Signature":     signature,
		}).
		SetMaxRetries(maxRetry).
		SetRetryWaitTime(retryWaitTime)

	resp, err := client.DoRequest(method, endpoint, body, queries)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %s", err.Error())
	}
	return resp, nil
}
