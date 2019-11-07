package dto

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type    string             `bson:"type" json:"type"`
	Capital int64              `bson:"capital" json:"capital"`
	Name    string             `bson:"name" json:"name"`
	UID     string             `bson:"id" json:"id"`
}

type CompanyParams struct {
	ID      *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type    *string             `bson:"type,omitempty" json:"type,omitempty"`
	Capital *int64              `bson:"capital,omitempty" json:"capital,omitempty"`
	Name    *string             `bson:"name,omitempty" json:"name,omitempty"`
	UID     *string             `bson:"id,omitempty" json:"id,omitempty"`
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
	Limit int64
	Head  primitive.ObjectID
}

func GetCompanies(db *mongo.Database, params *CompanyParams, opts Options) *[]Company {
	companies := []Company{}
	p := bson.M{}
	var err error

	options := options.Find()
	if opts.Limit != 0 {
		options.SetLimit(opts.Limit)
	}

	var defaultID primitive.ObjectID

	if opts.Head != defaultID {
		p["_id"] = bson.M{"$gt": opts.Head}
	}

	common.MergeStructToMap(params, p)
	fmt.Printf("p: %+v \n", p)

	cur, err := db.Collection("companies").Find(
		context.Background(),
		p,
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

	return &companies
}
