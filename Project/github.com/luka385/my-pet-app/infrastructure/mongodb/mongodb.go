package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/luka385/my-pet-app/domain"
)

type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoDB(connectionString, dbName, collectionName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	coll := db.Collection(collectionName)

	return &MongoDB{
		client:     client,
		database:   db,
		collection: coll,
	}, nil
}

func (m *MongoDB) Create(pet *domain.Pet) error {
	_, err := m.collection.InsertOne(context.Background(), pet)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetByID(id string) (*domain.Pet, error) {
	//objID, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return nil, err
	//}

	filter := bson.M{"id": id}
	var pet domain.Pet
	err := m.collection.FindOne(context.Background(), filter).Decode(&pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (m *MongoDB) Update(id string, updatePet *domain.Pet) error {
	//objID, err := primitive.ObjectIDFromHex(pet.ID)
	//if err != nil {
	//	return err
	//}

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatePet}

	_, err := m.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Delete(id string) error {
	//objID, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return err
	//}

	filter := bson.M{"id": id}
	_, err := m.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
