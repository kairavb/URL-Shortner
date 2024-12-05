package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index.html", fiber.Map{})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		url := c.FormValue("url")

		if url == "" {
			return c.Render("views/index.html", fiber.Map{"url": "URL cannot be empty"})
		}

		surl := GenerateShortURL()

		return c.Render("views/index.html", fiber.Map{"url": surl})
	})

	app.Listen(":5000")
}

func GenerateShortURL() string {
	surl := 6
	return surl
}
