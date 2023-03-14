package store

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	locaColl *mongo.Collection
}


func Connect() Store {
	godotenv.Load()
	uri := os.Getenv("mongouri")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("info")

	return Store{
		locaColl: db.Collection("patient"),
	}
}
