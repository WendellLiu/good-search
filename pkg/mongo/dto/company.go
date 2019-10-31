package dto

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetCompany(p *CompanyParams) *Company {
	company := Company{}
	var err error

	collection := mongo.DB.Collection("companies")
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

func GetCompanies(p *CompanyParams, limit int64) *[]Company {
	companies := []Company{}
	var err error
	options := options.Find()

	// Limit by 10 documents only
	options.SetLimit(limit)

	cur, err := mongo.DB.Collection("companies").Find(
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
