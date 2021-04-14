package api

import (
	"challenge.haraj.com.sa/kraicklist/services"
	"github.com/gofiber/fiber/v2"
)

type respWrapper struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponsePayload interface{} `json:"responsePayload,omitempty"`
}

type handler struct {
	app     *fiber.App
	service *services.Service
}

func Handler(app *fiber.App, service *services.Service) {
	handler := &handler{
		app:     app,
		service: service,
	}

	AdsDataHandler(handler)
}
