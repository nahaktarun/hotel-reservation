package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nahaktarun/hotel-reservation/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}
	fmt.Println(client)

	ListenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	log.Fatal(app.Listen(*ListenAddr))

}
