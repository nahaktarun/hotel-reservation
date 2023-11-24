package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Working just fine..."})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Tarun Nahak"})
}

func main() {

	app := fiber.New()

	apiV1 := app.Group("api/v1")
	app.Get("/foo", handleFoo)
	apiV1.Get("/user", handleUser)
	log.Fatal(app.Listen(":3000"))

}
