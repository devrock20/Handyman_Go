package connections

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://myuser:mypassword@cluster0.10yhq.mongodb.net/MyProject?retryWrites=true&w=majority")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, ctx, cancel
}
