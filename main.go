package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
	"github.com/rs/zerolog/log"

	"github.com/afourni/fizzbuzz/apierrors"
	"github.com/afourni/fizzbuzz/resource"
)

const version = "v1"

func main() {
	var bindAddress string

	flag.StringVar(&bindAddress, "bind-address", ":8080", "The bind address used by the FIZZBUZZ service")
	flag.Parse()

	r := gin.Default()

	// It builds the HTTP response using the errors pushed by the handlers
	r.Use(apierrors.ErrorMiddleware())

	// Handle custom 404 HTTP response
	r.NoRoute(func(c *gin.Context) {
		errNotFound := &apierrors.APIError{
			Messages: []string{"Resource not found"},
			Status:   http.StatusNotFound,
		}
		c.JSON(http.StatusNotFound, errNotFound)
	})

	api := r.Group("/api")
	v1 := api.Group(version)

	v1.GET(resource.FizzBuzzPath, resource.FizzBuzzHandler)

	err := r.Run(bindAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to run the API")
	}
}
