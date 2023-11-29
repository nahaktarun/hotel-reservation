package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nahaktarun/hotel-reservation/api"
	"github.com/nahaktarun/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "hotel-reservation"
const userCol1 = "users"

func main() {
	ListenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiV1 := app.Group("api/v1")

	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	log.Fatal(app.Listen(*ListenAddr))

}
