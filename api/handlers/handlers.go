package handlers

import (
    "github.com/gofiber/fiber/v2"
)

func GetDashboard(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "Dashboard data",
    })
}

func GetTickets(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "List of tickets",
    })
}

func CreateTicket(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "Ticket created",
    })
}
