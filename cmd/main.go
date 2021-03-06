package main

import (
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	"github.com/wendellliu/good-search/pkg/logger"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/queue"
	"github.com/wendellliu/good-search/pkg/queue/consumer"
	"github.com/wendellliu/good-search/pkg/server"
)

func main() {
	err := config.Load()
	logger.Load()

	if err != nil {
		logger.Logger.Fatal(err)
	}

	db, err := mongo.New()
	repository := &dto.Repository{DB: db}

	elasticsearch, err := es.New()

	queue, err := queue.New(&consumer.Dependencies{
		Es:   elasticsearch,
		Repo: repository,
	})

	if err != nil {
		logger.Logger.Fatal(err)
	}

	server.Load(server.Dependencies{
		Repo:  repository,
		Es:    elasticsearch,
		Queue: queue,
	})

	defer queue.Conn.Close()
}
