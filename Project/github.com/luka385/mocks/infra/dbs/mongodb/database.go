package dbs

import (
	"context"
	"primer-api/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*domain.User, error) {

	filter := bson.M{"id": id}

	user := &domain.User{}

	err := r.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
