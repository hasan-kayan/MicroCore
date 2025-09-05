package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientBundle struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

func Connect(uri, db, coll string) (*ClientBundle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := cl.Ping(ctx, nil); err != nil {
		return nil, err
	}

	d := cl.Database(db)
	c := d.Collection(coll)

	// index/validator eklemek istersen burada ekleyebilirsin

	return &ClientBundle{Client: cl, DB: d, Collection: c}, nil
}
