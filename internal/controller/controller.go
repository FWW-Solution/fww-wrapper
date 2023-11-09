package controller

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_passanger"
	"fww-wrapper/internal/tools"
	"strconv"

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

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		return ctx.JSON(errValidation)
	}

	result, err := c.adapter.RegisterPassanger(&body)
	if err != nil {
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.JSON(response)
}

func (c *Controller) DetailPassanger(ctx *fiber.Ctx) error {
	data := ctx.Query("data")
	dataInt, err := strconv.Atoi(data)

	if err != nil {
		return err
	}

	result, err := c.adapter.GetPassanger(dataInt)
	if err != nil {
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.JSON(response)
}

func (c *Controller) UpdatePassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestUpdate

	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		return ctx.JSON(errValidation)
	}

	result, err := c.adapter.UpdatePassanger(&body)
	if err != nil {
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.JSON(response)

}
