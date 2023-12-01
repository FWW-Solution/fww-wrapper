package controller

import (
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_user"
	"fww-wrapper/internal/tools"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var body dto_user.RequestLogin

	if err := ctx.BodyParser(&body); err != nil {
		c.Log.Error(err)
		return err
	}

	result, err := c.UseCase.GenerateToken(body.Username, body.Password)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	data := dto_user.ResponseLogin{
		Token: result,
	}

	response := tools.ResponseBuilder(data, meta)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
