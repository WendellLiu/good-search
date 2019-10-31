package mongo

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

func init() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected")
}
