package dto

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	cur := mongo.DB.Collection("companies").FindOne(
		context.Background(),
		p,
	)
	err = cur.Decode(&company)
	if err != nil {
		fmt.Println(err)
	}

	return &company
}
