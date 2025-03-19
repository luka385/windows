package dbs

import (
	"Project/github.com/luka385/mocks2/domain"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (ur *UserRepository) GetUserById(id string) (*domain.User, error) {
	filter := bson.M{"id": id}
	var user domain.User

	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User

	cursor, err := ur.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		copieduser := user
		users = append(users, &copieduser)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user *domain.User) error {
	_, err := ur.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateUser(id string, user *domain.User) error {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"email":    user.Email,
			"password": user.Password,
		},
	}

	_, err := ur.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(id string) error {
	filter := bson.M{"id": id}

	_, err := ur.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
