package dbs

import (
	"context"
	"primer-api/domain"

	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) GetById(id string) (*domain.User, error) {
	filter := bson.M{"id": id}

	user := &domain.User{}

	err := r.collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return err
}
