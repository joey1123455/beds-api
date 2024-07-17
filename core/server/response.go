package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/internal/validator"
)

// Response represents a standard JSON response structure.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// SendSuccess sends a JSON response with a status code of 200 (OK).
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusText(http.StatusOK),
		Message: message,
		Data:    data,
	})
}

// SendCreated sends a JSON response with a status code of 201 (Created).
func SendCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  http.StatusText(http.StatusCreated),
		Message: message,
		Data:    data,
	})
}

// SendError sends a JSON response with a specified status code and error message.
func SendError(c *gin.Context, statusCode int, errs ...error) {

	responseData := Response{
		Status:  http.StatusText(statusCode),
		Message: "error processing request",
	}

	outputErrors := make([]string, 0, len(errs))
	for _, err := range errs {
		outputErrors = append(outputErrors, err.Error())
	}
	responseData.Errors = outputErrors

	c.JSON(statusCode, responseData)
}

// SendBadRequest sends a JSON response with a status code of 400 (Bad Request).
func SendBadRequest(c *gin.Context, err error) {
	SendError(c, http.StatusBadRequest, err)
}

// SendForbidden sends a JSON response with a status code of 403 (Forbidden).
func SendForbidden(c *gin.Context, err error) {
	SendError(c, http.StatusForbidden, err)
}

// SendInternalServerError sends a JSON response with a status code of 500 (Internal Server Error).
func SendInternalServerError(c *gin.Context, err error) {
	SendError(c, http.StatusInternalServerError, err)
}

// SendMethodNotAllowedError sends a JSON response with a status code of 405 (Method Not Allowed).
func SendMethodNotAllowedError(c *gin.Context, err error) {
	SendError(c, http.StatusMethodNotAllowed, err)
}

// SendNotFound sends a JSON response with a status code of 404 (Not Found).
func SendNotFound(c *gin.Context, err error) {
	SendError(c, http.StatusNotFound, err)
}

// SendConflict sends a JSON response with a status code of 409 (Unauthorized).
func SendConflict(c *gin.Context, err error) {
	SendError(c, http.StatusConflict, err)
}

// SendUnauthorized sends a JSON response with a status code of 401 (Unauthorized).
func SendUnauthorized(c *gin.Context, err error) {
	SendError(c, http.StatusUnauthorized, err)
}

func SendValidationError(c *gin.Context, errors *validator.ValidationError) {
	responseData := Response{
		Errors:  errors.Fields,
		Message: errors.Message,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
	}

	c.JSON(http.StatusUnprocessableEntity, responseData)

}

func SendParsingError(c *gin.Context, err error) {
	SendError(c, http.StatusUnprocessableEntity, err)
}
