package resource

import (
	"io"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/afourni/fizzbuzz/apierrors"
)

// FizzBuzzPath is the relative path used to handle the FizzBuzz query
const FizzBuzzPath = "/fizzbuzz"

// ErrInvalidParameters is the error returned in case of invalid parameters
const ErrInvalidParameters apierrors.Error = "Invalid parameters"

// FizzBuzzParams represents all the query parameters used to handle the FizzBuzz query
type FizzBuzzParams struct {
	Int1    int64  `form:"int1" binding:"required,min=1"`
	String1 string `form:"string1" binding:"required"`
	Int2    int64  `form:"int2" binding:"required,min=1"`
	String2 string `form:"string2" binding:"required"`
	Limit   int64  `form:"limit" binding:"required,min=1"`
}

// fizzBuzz is the business function which compute the FizzBuzz
// The FizzBuzz is processed as a stream to reduce memory allocation
func fizzBuzz(fizzBuzzParams FizzBuzzParams, w io.Writer) error {

	// Check the input parameters validity
	if fizzBuzzParams.Int1 <= 0 || fizzBuzzParams.Int2 <= 0 || fizzBuzzParams.Limit <= 0 {
		return ErrInvalidParameters
	}

	if _, err := w.Write([]byte(`["`)); err != nil {
		return err
	}

	for i := int64(1); i <= fizzBuzzParams.Limit; i++ {
		if i > 1 {
			if _, err := w.Write([]byte(`,"`)); err != nil {
				return err
			}
		}

		moduloInt1 := i%fizzBuzzParams.Int1 == 0
		moduloInt2 := i%fizzBuzzParams.Int2 == 0

		var value string
		if moduloInt1 || moduloInt2 {
			if moduloInt1 {
				value = fizzBuzzParams.String1
			}
			if moduloInt2 {
				value += fizzBuzzParams.String2
			}
		} else {
			value = strconv.FormatInt(i, 10)
		}

		if _, err := w.Write([]byte(value + `"`)); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte("]")); err != nil {
		return err
	}
	return nil
}

// FizzBuzzHandler is the handler used to compute the FizzBuzz
func FizzBuzzHandler(c *gin.Context) {
	var fizzBuzzParams FizzBuzzParams

	// Check the input parameters validity
	if err := c.ShouldBind(&fizzBuzzParams); err != nil {
		c.Error(err)
		return
	}

	c.Header("Content-Type", "application/json")

	if err := fizzBuzz(fizzBuzzParams, c.Writer); err != nil {
		c.Error(err)
		return
	}
}
