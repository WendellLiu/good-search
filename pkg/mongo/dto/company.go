package dto

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultLimit = 100
)

type Company struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type    string             `bson:"type" json:"type"`
	Capital int                `bson:"capital" json:"capital"`
	Name    string             `bson:"name" json:"name"`
	UID     string             `bson:"id" json:"id"`
}

type CompanyParams struct {
	Type    *string `bson:"type,omitempty" json:"type,omitempty"`
	Capital *int    `bson:"capital,omitempty" json:"capital,omitempty"`
	Name    *string `bson:"name,omitempty" json:"name,omitempty"`
}

func GetCompany(db *mongo.Database, p *CompanyParams) *Company {
	company := Company{}
	var err error

	collection := db.Collection("companies")
	cur := collection.FindOne(
		context.Background(),
		p,
	)
	err = cur.Decode(&company)
	if err != nil {
		fmt.Println(err)
	}

	return &company
}

type Options struct {
	Limit    int64
	CursorID primitive.ObjectID
}

func GetCompanies(db *mongo.Database, params *CompanyParams, opts Options) []Company {
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
	fmt.Printf("query: %+v \n", query)
	fmt.Printf("params capital %d \n", params.Capital)
	fmt.Printf("capital %d \n", query["capital"])
	cur, err := db.Collection("companies").Find(
		context.Background(),
		query,
		options,
	)
	defer cur.Close(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	err = cur.All(context.Background(), &companies)
	if err != nil {
		fmt.Println(err)
	}

	return companies
}
