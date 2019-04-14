package apierrors

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"gopkg.in/go-playground/validator.v8"
)

// Error is the type of errors used in the business code
type Error string

// Error is the method required to implement the interface error
func (e Error) Error() string { return string(e) }

// APIError contains all the data exposed by the API in case of error
type APIError struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
}

// ErrorMiddleware is the handler called after each business handler
// It builds the HTTP response depending on the returned error
func ErrorMiddleware() gin.HandlerFunc {
	return errorMiddleware(gin.ErrorTypeAny)
}

func errorMiddleware(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			apiError := &APIError{
				Status: http.StatusInternalServerError,
			}

			switch typedError := detectedErrors[0].Err.(type) {

			case validator.ValidationErrors: // An error during the validation of the input parameters
				apiError.Status = http.StatusBadRequest
				for _, fieldError := range typedError {
					log.Warn().Err(typedError).Msg("Invalid parameters")

					apiError.Messages = append(apiError.Messages, fmt.Sprintf("Invalid value for parameter '%s'", fieldError.Name))
				}

			case *strconv.NumError: // An error during the parsing of an input parameter
				log.Warn().Err(typedError).Msg("Parameters parsing error")

				apiError.Status = http.StatusBadRequest
				apiError.Messages = []string{"Parameters parsing error"}

			default: // Unhandled error: It returns 500
				log.Error().Err(typedError).Msg("Internal server error")

				apiError.Messages = []string{"Internal server error"}
			}

			c.AbortWithStatusJSON(apiError.Status, apiError)
		}
	}
}
