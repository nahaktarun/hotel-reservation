package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Working just fine..."})
}

func main() {

	app := fiber.New()

	
	app.Get("/foo", handleFoo)
	log.Fatal(app.Listen(":3000"))

}
