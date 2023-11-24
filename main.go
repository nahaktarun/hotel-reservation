package main

import (
	"flag"
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

	ListenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")

	app := fiber.New()

	apiV1 := app.Group("api/v1")
	app.Get("/foo", handleFoo)
	apiV1.Get("/user", handleUser)
	log.Fatal(app.Listen(*ListenAddr))

}
