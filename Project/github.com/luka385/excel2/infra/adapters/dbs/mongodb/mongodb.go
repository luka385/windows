package dbs

import (
	"context"

	"github.com/luka385/excel2/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PersonRepository struct {
	collection *mongo.Collection
}

func NewPersonRepository(db *mongo.Database) *PersonRepository {
	return &PersonRepository{
		collection: db.Collection("persons"),
	}
}

func (r *PersonRepository) SavePerson(ctx context.Context, persons []domain.Person) ([]string, error) {
	var docs []interface{}

	for _, p := range persons {
		docs = append(docs, bson.M{
			"first_name": p.FirstName,
			"last_name":  p.LastName,
			"age":        p.Age,
			"phone":      p.Phone,
		})
	}

	_, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
