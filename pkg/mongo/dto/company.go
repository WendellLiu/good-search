package dto

import (
	"context"
	"fmt"

	"github.com/wendellliu/good-search/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type     string             `bson:"type" json:"type"`
	Capital  int64              `bson:"capital" json:"capital"`
	Name     string             `bson:"name" json:"name"`
	IDinItem string             `bson:"id" json:"id"`
}

func GetCompany() *Company {
	company := Company{}
	cur := mongo.DB.Collection("companies").FindOne(context.Background(), bson.M{
		"name": "復華廣告有限公司",
	})
	err := cur.Decode(&company)
	if err != nil {
		fmt.Println(err)
	}

	return &company
}
