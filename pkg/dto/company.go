package dto

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Company struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type    string             `bson:"type" json:"type"`
	Capital int                `bson:"capital" json:"capital"`
	Name    string             `bson:"name" json:"name"`
	UID     string             `bson:"id" json:"id"`
}

type CompaniesParams struct {
	Type    *string `bson:"type,omitempty" json:"type,omitempty"`
	Capital *int    `bson:"capital,omitempty" json:"capital,omitempty"`
	Name    *string `bson:"name,omitempty" json:"name,omitempty"`
}

func GetCompany(db dbAdapter.Database, id string) (*Company, error) {
	collectionName := "companies"

	var err error

	collection := db.UseTable(collectionName)

	result, err := collection.QueryOne(
		context.Background(),
		id,
	)

	fmt.Printf("result: %+v \n", result)

	value, ok := result.(Company)
	fmt.Printf("result type = %T \n", result)
	fmt.Printf("value type = %T \n", value)
	if !ok {
		err = fmt.Errorf("type error: %T", value)
	}

	return &value, err

}

func GetCompanies(db *mongo.Database, params *CompaniesParams, opts Options) []Company {
	collectionName := "companies"
	companies := []Company{}
	query := bson.M{}
	var err error

	options := options.Find()
	if opts.Limit != 0 {
		options.SetLimit(opts.Limit)
	} else {
		options.SetLimit(defaultLimit)

	}
	var defaultID primitive.ObjectID

	if opts.CursorID != defaultID {
		query["_id"] = bson.M{
			"$gt": opts.CursorID,
		}
	}

	if params.Type != nil {
		query["type"] = params.Type
	}
	if params.Capital != nil {
		query["capital"] = params.Capital
	}
	if params.Name != nil {
		query["name"] = params.Name
	}

	cur, err := db.Collection(collectionName).Find(
		context.Background(),
		query,
		options,
	)
	defer cur.Close(context.Background())
	if err != nil {
		logger.Logger.Error(err)
	}

	err = cur.All(context.Background(), &companies)
	if err != nil {
		logger.Logger.Error(err)
	}

	return companies
}