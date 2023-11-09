package controller

import (
	"fww-wrapper/internal/adapter"
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_passanger"
	"fww-wrapper/internal/tools"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	Adapter adapter.Adapter
	Log     *zap.SugaredLogger
}

func (c *Controller) RegisterPassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestRegister

	if err := ctx.BodyParser(&body); err != nil {
		c.Log.Error(err)
		return err
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		c.Log.Error(errValidation)
		return ctx.JSON(errValidation)
	}

	result, err := c.Adapter.RegisterPassanger(&body)
	if err != nil {
		c.Log.Error(err)
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
		c.Log.Error(err)
		return err
	}

	result, err := c.Adapter.GetPassanger(dataInt)
	if err != nil {
		c.Log.Error(err)
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
		c.Log.Error(err)
		return err
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		c.Log.Error(errValidation)
		return ctx.JSON(errValidation)
	}

	result, err := c.Adapter.UpdatePassanger(&body)
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
