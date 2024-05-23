package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang_mongoDB/controllers"
)

func main() {
	r := gin.Default()
	uc := controllers.NewUserController(getMongoDBClient())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/user/:id", uc.UpdateUser)

	r.Run(":9000") // 默认情况下，gin.Run 会使用 http.ListenAndServe
}

func getMongoDBClient() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Ping the MongoDB server to confirm a successful connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
