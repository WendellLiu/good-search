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

const (
	defaultLimit = 10
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

func (collection *MongoCollection) QueryOne(ctx context.Context, id string, result interface{}) error {
	var err error
	ID, err := primitive.ObjectIDFromHex(id)
	cur := collection.collection.FindOne(
		ctx,
		bson.M{"_id": ID},
	)
	err = cur.Decode(result)
	return err
}

func (collection *MongoCollection) QueryPagination(
	ctx context.Context,
	params map[string]interface{},
	opts dbAdapter.Options,
	results interface{},
) error {
	//localLogger := logger.Logger.WithFields(
	//logrus.Fields{"endpoint": "mongo-QueryPagination"},
	//)
	query := bson.M{}

	options := options.Find()
	if opts.Limit != 0 {
		options.SetLimit(opts.Limit)
	} else {
		options.SetLimit(defaultLimit)

	}
	var defaultID primitive.ObjectID
	cursorID, encodeError := primitive.ObjectIDFromHex(opts.CursorID)

	if encodeError != nil {
		return fmt.Errorf("%s - id: %+v", encodeError, opts.CursorID)
	}

	if cursorID != defaultID {
		query["_id"] = bson.M{
			"$gt": cursorID,
		}
	}

	for k, v := range params {
		if v != nil {
			query[k] = v
		}
	}

	//localLogger.Infof("query: %+v \n", query)
	//localLogger.Infof("options: %+v \n", options)

	cur, cursorBuildingError := collection.collection.Find(
		context.Background(),
		query,
		options,
	)

	if cursorBuildingError != nil {
		return fmt.Errorf("cursor building error: %s", cursorBuildingError)
	}
	defer cur.Close(context.Background())

	cursorRunErr := cur.All(context.Background(), results)

	if cursorRunErr != nil {
		return fmt.Errorf("collection cursor run error: %s", cursorRunErr)
	}
	return nil
}
