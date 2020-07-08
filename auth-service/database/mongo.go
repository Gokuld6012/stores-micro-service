package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConn returns a mongo connection pool.
func DBConn() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		// need to check whether this format is correct or not.
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database("micro-auth-service")
	return db, nil
}
