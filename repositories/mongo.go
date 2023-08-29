package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hexagonal-fiber-impl/core/domain"
	"hexagonal-fiber-impl/core/ports"
	"time"
)

const (
	MongoClientTimeout = 20
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(conn string) *UserRepository {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}
	return &UserRepository{
		client:     client,
		database:   client.Database("hexagonal"),
		collection: client.Database("hexagonal").Collection("hexagonal"),
	}
}

func (r *UserRepository) Login(email string, password string) (domain.User, error) {
	var user domain.User
	filter := bson.D{{"email", email}}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Register(email string, password string) (domain.User, error) {
	//Here your code for save in mongo database
	newUser := domain.User{Email: email, Password: password, Id: primitive.NewObjectID()}
	_, err := r.collection.InsertOne(context.TODO(), &newUser)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, err
}
