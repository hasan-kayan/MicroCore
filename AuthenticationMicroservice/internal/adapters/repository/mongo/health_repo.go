package mongo

import (
	"context"

	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthRepo struct {
	col Collection
}

type Collection interface {
	InsertOne(ctx context.Context, document any, opts ...any) (any, error)
}

func NewHealthRepo(cb *ClientBundle) *HealthRepo {
	return &HealthRepo{col: cb.Collection}
}

func (r *HealthRepo) Insert(ctx context.Context, ev domain.HealthEvent) (string, error) {
	res, err := r.col.InsertOne(ctx, ev)
	if err != nil {
		return "", err
	}
	switch v := res.(type) {
	case *struct{ InsertedID any }:
		if id, ok := v.InsertedID.(primitive.ObjectID); ok {
			return id.Hex(), nil
		}
	}
	// generic path (mongo-go-driver gerçek dönen tipleri sarıyor)
	m, ok := res.(interface{ GetInsertedID() any })
	if ok {
		if oid, ok := m.GetInsertedID().(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
	}
	// son çare: BSON çöz
	b, _ := bson.Marshal(res)
	_ = b
	return "", nil
}
