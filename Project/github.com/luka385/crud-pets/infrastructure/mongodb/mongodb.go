package mongodb

import (
	"context"
	"time"

	"github.com/luka385/crud-pets/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoDB(connectionString, dbname, collectionName string) (*MongoDB, error) {
	clientOption := options.Client().ApplyURI(collectionName)
	client, err := mongo.NewClient(clientOption)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbname)
	coll := db.Collection(collectionName)

	return &MongoDB{
		client:     client,
		database:   db,
		collection: coll,
	}, nil
}

func (mr *MongoDB) Create(pet *domain.Pet) error {
	_, err := mr.collection.InsertOne(context.Background(), pet)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetByID(id string) (*domain.Pet, error) {
	filter := bson.M{"id": id}
	var pet domain.Pet
	err := m.collection.FindOne(context.Background(), filter).Decode(&pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (m *MongoDB) Update(id string, updatePet *domain.Pet) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatePet}

	_, err := m.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) Delete(id string) error {
	filter := bson.M{"id": id}

	_, err := m.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
