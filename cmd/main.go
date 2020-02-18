package main

import (
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	"github.com/wendellliu/good-search/pkg/logger"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/server"
)

func main() {
	logger.Load()
	config.Load()

	db, err := mongo.New()
	repository := &dto.Repository{DB: db}

	elasticsearch, err := es.New()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	server.Load(server.Dependencies{
		Repo: repository,
		Es:   elasticsearch,
	})
}
