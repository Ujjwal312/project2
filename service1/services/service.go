package services

import (
	"context"

	"example.com/m/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Orderservice interface {
	CreateOrder(*models.Order) error
}

type OrderserviceImpl struct {
	orderCollection *mongo.Collection
	Ctx             context.Context
}

func NeworderService(orderCollection *mongo.Collection, ctx context.Context) Orderservice {
	return &OrderserviceImpl{
		orderCollection: orderCollection,
		Ctx:             ctx,
	}
}

func (u *OrderserviceImpl) CreateOrder(order *models.Order) error {
	_, err := u.orderCollection.InsertOne(u.Ctx, order)
	return err
}
