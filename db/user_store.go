package db

import (
	"context"
	"fmt"

	"github.com/nahaktarun/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCol1 = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
	Dropper
}

type MongoUserStore struct {
	client *mongo.Client
	col1   *mongo.Collection
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("---- dropping user collection ---")
	return s.col1.Drop(ctx)
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {

	return &MongoUserStore{
		client: client,
		col1:   client.Database(dbname).Collection(userCol1),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User

	if err := s.col1.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {

	cur, err := s.col1.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}
	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.col1.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// TODO: May be its a good idea to handle if we did not delete any user
	// maybe log it or something
	_, err = s.col1.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {

	update := bson.D{
		{
			"$set", params.ToBSON(),
		},
	}
	_, err := s.col1.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
