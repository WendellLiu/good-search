package mongo

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/config"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	connection *mongo.Database
}

type MongoCollection struct {
	collection *mongo.Collection
}

// New mongo db
func New() (*MongoDB, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s", config.Config.MongoDBHost, config.Config.MongoDBPort)
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongoURI),
	)

	err = client.Ping(context.Background(), nil)
	conn := client.Database(config.Config.MongoDBName)
	return &MongoDB{connection: conn}, err
}

func (db *MongoDB) UseTable(collectionName string) dbAdapter.Table {
	return &MongoCollection{collection: db.connection.Collection(collectionName)}
}

func (collection *MongoCollection) QueryOne(ctx context.Context, id string) (interface{}, error) {
	var result interface{}
	var err error
	ID, err := primitive.ObjectIDFromHex(id)
	cur := collection.collection.FindOne(
		ctx,
		bson.M{"_id": ID},
	)
	err = cur.Decode(&result)
	return result, err
}

//func Load() {
//var err error
//client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
//if err != nil {
//fmt.Println(err)
//}
//err = client.Ping(context.Background(), nil)
//if err != nil {
//fmt.Println(err)
//}

//DB = client.Database(config.Config.MongoDBName)
//}
