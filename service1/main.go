package main

import (
	"context"
	"fmt"
	"log"

	controller "example.com/m/controllers"

	"example.com/m/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server          *gin.Engine
	orderservice    services.Orderservice
	ordercontroller controller.OrderController
	ctx             context.Context
	ordercollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	ordercollection = mongoclient.Database("orderdb").Collection("order")
	orderservice = services.NeworderService(ordercollection, ctx)
	ordercontroller = controller.New(orderservice)
	server = gin.Default()
}

// v1/order/create
func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	ordercontroller.RegisterorderRoutes(basepath)
	log.Fatal(server.Run(":9091"))
}
