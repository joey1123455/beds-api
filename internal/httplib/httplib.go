package httplib

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/valyala/fastjson"
	"gopkg.in/yaml.v2"
)

// HttpClient is a struct that holds the HTTP client, timeout settings, and retry configuration
type HttpClient struct {
	Client        *http.Client
	Timeout       time.Duration
	Headers       map[string]string
	MaxRetries    int
	RetryWaitTime time.Duration
}

// NewHttpClient creates a new HttpClient with default configuration
func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{Client: &http.Client{Timeout: timeout}, Timeout: timeout}
}

// SetHeaders sets the custom headers for the request and returns the client
func (c *HttpClient) SetHeaders(headers map[string]string) *HttpClient {
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
	for key, value := range headers {
		c.Headers[key] = value
	}
	return c
}

// SetMaxRetries sets the maximum number of retries for failed requests
func (c *HttpClient) SetMaxRetries(retries int) *HttpClient {
	c.MaxRetries = retries
	return c
}

// SetRetryWaitTime sets the wait time between retries
func (c *HttpClient) SetRetryWaitTime(waitTime time.Duration) *HttpClient {
	c.RetryWaitTime = waitTime
	return c
}

func buildURL(baseURL string, params map[string]string) (string, error) {
	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Create a query object from the base URL's query string
	query := u.Query()

	// Add the query parameters from the map
	for key, value := range params {
		query.Add(key, value)
	}

	// Encode the query parameters and set it to the URL
	u.RawQuery = query.Encode()

	// Return the full URL as a string
	return u.String(), nil
}

// DoRequest sends an HTTP request using the specified method, url, and interface for body
func (c *HttpClient) DoRequest(method, url string, body interface{}, queries map[string]string) (*http.Response, error) {
	var reqBody []byte
	var err error

	if queries != nil {
		url, err = buildURL(url, queries)
		if err != nil {
			return nil, err
		}
	}

	switch body.(type) {
	case []byte:
		reqBody = body.([]byte)
	default:
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i <= c.MaxRetries; i++ {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		// Set custom headers
		if c.Headers != nil {
			for key, value := range c.Headers {
				req.Header.Set(key, value)
			}
		}

		resp, err := c.Client.Do(req)
		if err != nil {
			return nil, err
		}
		log.Println(resp.StatusCode)
		// log.Println(resp.)
		if resp.StatusCode >= 200 && resp.StatusCode < 500 {
			return resp, nil
		}

		if i < c.MaxRetries {
			time.Sleep(c.RetryWaitTime)
		}
		return resp, nil
	}

	return nil, fmt.Errorf("request failed after %d retries", c.MaxRetries)
}

// ReadJsonResponse unmarshals the response body into the provided interface
func ReadJsonResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	var reader io.ReadCloser
	var err error

	// Check if the response is gzip-compressed
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
	} else {
		reader = resp.Body
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data
	return json.Unmarshal(bodyBytes, v)
}

func ReadYAMLResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bodyBytes, v)
}

func ReadRawBodyResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var p fastjson.Parser
	z, err := p.ParseBytes(bodyBytes)
	if err != nil {
		return err
	}

	validJSON := z.String()
	if err != nil {
		return err
	}

	log.Println(validJSON)

	// return json.Unmarshal(validJSON, v)
	return nil
}

func ParseCustomFormat(input string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	var currentSection string

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if !strings.Contains(line, ":") {
			currentSection = line
			result[currentSection] = make(map[string]string)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			value = strings.Trim(value, "'")
			result[currentSection][key] = value
		}
	}

	return result
}
