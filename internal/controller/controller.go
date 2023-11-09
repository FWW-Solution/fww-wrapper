package controller

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_passanger"
	"fww-wrapper/internal/tools"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	adapter adapter.Adapter
}

func (c *Controller) RegisterPassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestRegister

	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	result, err := c.adapter.RegisterPassanger(&body)
	if err != nil {
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.JSON(response)
}

func (c *Controller) DetailPassanger(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) UpdatePassanger(ctx *fiber.Ctx) error {
	return nil
}
