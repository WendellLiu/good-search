package main

import (
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	"github.com/wendellliu/good-search/pkg/logger"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file")
	}

	logger.Load()
	config.Load()

	db, err := mongo.New()
	repository := &dto.Repository{DB: db}

	elasticsearch, err := es.New()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	server.Load(repository, elasticsearch)
}
