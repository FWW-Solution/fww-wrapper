package router

import (
	"fww-wrapper/internal/controller"
	"fww-wrapper/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Initialize(app *fiber.App, ctrl *controller.Controller, m *middleware.Middleware) *fiber.App {
	app.Get("/", monitor.New(monitor.Config{Title: "fww-wrapper metrics page"}))

	wrapper := app.Group("/wrapper")

	Api := wrapper.Group("/api")

	v1 := Api.Group("/v1")

	// User
	v1.Post("/user/login", ctrl.Login)

	// Passanger
	v1.Post("/passanger", m.ValidateAPIKey, ctrl.RegisterPassanger)
	v1.Get("/passanger", m.ValidateAPIKey, ctrl.DetailPassanger)
	v1.Put("/passanger", m.ValidateAPIKey, ctrl.UpdatePassanger)

	// Airport
	v1.Get("/airports", m.ValidateJWTUser, ctrl.GetAirport)

	//Flight
	v1.Get("/flights", m.ValidateAPIKey, ctrl.GetFlights)
	v1.Get("/flight", m.ValidateAPIKey, ctrl.GetDetailFlightByID)

	// Booking
	v1.Post("/booking", m.ValidateJWTUser, ctrl.Booking)
	v1.Get("/booking", m.ValidateJWTUser, ctrl.GetDetailBooking)

	// Payment
	v1.Post("/payment", m.ValidateJWTUser, ctrl.DoPayment)
	// Payment
	v1.Get("/payment/status", m.ValidateJWTUser, ctrl.GetPaymentStatus)
	v1.Get("/payment/methods", m.ValidateJWTUser, ctrl.GetPaymentMethods)

	// ticket
	v1.Post("/ticket/redeem", m.ValidateJWTUser, ctrl.RedeemTicket)
	return app
}
