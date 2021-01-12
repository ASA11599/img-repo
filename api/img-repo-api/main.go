package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func getImage(c * fiber.Ctx) error {
	return errors.New("An error")
}

func main() {
	app := fiber.New()
	app.Get("/", getImage)
	log.Fatal(app.Listen("0.0.0.0:80"))
}
