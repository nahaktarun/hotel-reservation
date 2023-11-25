package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nahaktarun/hotel-reservation/api"
	"github.com/nahaktarun/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "hotel-reservation"
const userCol1 = "users"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	col1 := client.Database(dbName).Collection(userCol1)

	user := types.User{
		FirstName: "Tarun",
		LastName:  "Nahak",
	}
	res, err := col1.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	ListenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	log.Fatal(app.Listen(*ListenAddr))

}
