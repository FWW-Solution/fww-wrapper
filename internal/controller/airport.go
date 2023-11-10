package controller

import (
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/tools"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAirport(ctx *fiber.Ctx) error {
	city := ctx.Query("city", "")
	province := ctx.Query("province", "")
	iata := ctx.Query("iata", "")

	result, err := c.Adapter.GetAirport(city, province, iata)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}
