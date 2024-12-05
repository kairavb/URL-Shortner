package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index.html", fiber.Map{})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.Render("views/index.html", fiber.Map{"url": "answer"})
	})

	app.Listen(":5000")
}
