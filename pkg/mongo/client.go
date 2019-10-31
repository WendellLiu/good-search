package mongo

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/config"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database

func Load() {
	mongoURI := fmt.Sprintf("mongodb://%s:%s", config.Config.MongoDBHost, config.Config.MongoDBPort)
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	DB = client.Database(config.Config.MongoDBName)
}
