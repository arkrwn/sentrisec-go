package routes

import (
    "github.com/gofiber/fiber/v2"
    "sentrisec/api/handlers"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")

    api.Get("/dashboard", handlers.GetDashboard)
    api.Get("/tickets", handlers.GetTickets)
    api.Post("/tickets", handlers.CreateTicket)
    // Add more routes as needed
}
