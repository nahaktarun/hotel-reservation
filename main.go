package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nahaktarun/hotel-reservation/api"
)

func main() {

	ListenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	apiV1 := app.Group("api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	log.Fatal(app.Listen(*ListenAddr))

}
