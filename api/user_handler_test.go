package api

import (
	"context"
	"log"
	"testing"

	"github.com/nahaktarun/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testMongoUri = "mongodb://localhost:27017"
const dbname = "hotel-reservation-test"

type testDb struct {
	db.UserStore
}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testMongoUri))
	if err != nil {
		log.Fatal(err)
	}
	return &testDb{UserStore: db.NewMongoUserStore(client, dbname)}
}

func (tdb *testDb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
}
