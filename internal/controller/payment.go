package controller

import (
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_payment"
	"fww-wrapper/internal/tools"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) DoPayment(ctx *fiber.Ctx) error {
	var body dto_payment.Request

	if err := ctx.BodyParser(&body); err != nil {
		c.Log.Error(err)
		response := tools.ResponseBadRequest(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	paymentCode, err := c.Adapter.DoPayment(&body)
	if err != nil {
		c.Log.Error(err)
		response := tools.ResponseInternalServerError(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	meta := dto.MetaResponse{
		StatusCode: "OK200",
		IsSuccess:  true,
		Message:    "Payment success",
	}

	response := tools.ResponseBuilder(paymentCode, meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) GetPaymentStatus(ctx *fiber.Ctx) error {
	paymentCode := ctx.Query("payment_code", "")
	result, err := c.Adapter.GetPaymentStatus(paymentCode)
	if err != nil {
		c.Log.Error(err)
		response := tools.ResponseInternalServerError(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	meta := dto.MetaResponse{
		StatusCode: "OK200",
		IsSuccess:  true,
		Message:    "Payment success",
	}

	response := tools.ResponseBuilder(result, meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) GetPaymentMethods(ctx *fiber.Ctx) error {
	result, err := c.Adapter.GetPaymentMethods()
	if err != nil {
		c.Log.Error(err)
		response := tools.ResponseInternalServerError(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	meta := dto.MetaResponse{
		StatusCode: "OK200",
		IsSuccess:  true,
		Message:    "Payment success",
	}

	response := tools.ResponseBuilder(result, meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
