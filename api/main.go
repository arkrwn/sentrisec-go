package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "sentrisec/routes"
)

func main() {
    app := fiber.New()

    // Middleware
    app.Use(logger.New())

    // Serve static files from the 'dist' and 'static' directories
    app.Static("/dist", "../dashboard/dist")
    app.Static("/static", "../dashboard/static")

    // Dynamic route for all pages in the 'pages' directory
    app.Get("/:page", func(c *fiber.Ctx) error {
        page := c.Params("page")
        return c.SendFile("../dashboard/" + page)
    })

    // Default route to serve the index page
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("../dashboard/index.html")
    })

    // Setup API routes
    routes.Setup(app)

    // Start server
    log.Fatal(app.Listen(":3000"))
}
