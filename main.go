package main

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length = 6 // Length of the short url

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
	result := make([]byte, length) // Create an empty slice with zeros

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // Randomly pick a character from the letters string
	}

	return string(result)
}
