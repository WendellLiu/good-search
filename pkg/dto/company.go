package dto

import (
	"context"

	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetCompany(ctx context.Context, db dbAdapter.Database, id string) (Company, error) {
	collectionName := "companies"
	result := Company{}
	var err error

	collection := db.UseTable(collectionName)

	err = collection.QueryOne(
		context.Background(),
		id,
		&result,
	)

	if err != nil {
		logger.Logger.Error(err)
	}

	return result, err

}

func GetCompanies(
	ctx context.Context,
	db dbAdapter.Database,
	params *CompaniesParams,
	opts dbAdapter.Options,
) (companies []Company, err error) {
	collectionName := "companies"
	results := []Company{}

	query := make(map[string]interface{})
	if params.Type != nil {
		query["type"] = params.Type
	}
	if params.Capital != nil {
		query["capital"] = params.Capital
	}
	if params.Name != nil {
		query["name"] = params.Name
	}
	collection := db.UseTable(collectionName)
	err = collection.QueryPagination(
		ctx,
		query,
		opts,
		&results,
	)

	if err != nil {
		logger.Logger.Error(err)
	}

	return results, err
}
