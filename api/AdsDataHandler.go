package api

import "github.com/gofiber/fiber/v2"

func AdsDataHandler(handler *handler) {
	handler.app.Get("/search", handler.handleSearchAdsData)
}

func (handler *handler) handleSearchAdsData(c *fiber.Ctx) error {
	ctx := c.Context()

	searchStr := c.Query("searchStr")
	var resp respWrapper
	adsDatas, err := handler.service.AdsDataService.SearchAdsData(ctx, searchStr)
	if err != nil {
		resp.ResponseCode = fiber.StatusInternalServerError
		resp.ResponseMessage = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	resp.ResponseCode = fiber.StatusOK
	resp.ResponseMessage = "SUCCESS"
	resp.ResponsePayload = adsDatas
	return c.Status(fiber.StatusOK).JSON(resp)
}
