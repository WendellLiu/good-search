package consumer

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	"github.com/wendellliu/good-search/pkg/logger"
)

type Dependencies struct {
	Repo dto.DTO
	Es   es.Elasticsearch
}

func UpdateExperienceConsumer(dep *Dependencies, d *amqp.Delivery) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "UpdateExperienceConsumer"},
	)

	fmt.Printf("Received an experience id to be update: %s \n", d.Body)

	experience, err := dep.Repo.GetExperience(context.Background(), string(d.Body))

	if err != nil {
		localLogger.Error(err)
	}

	err = dep.Es.IndexExperience(context.Background(), experience)

	if err != nil {
		localLogger.Error(err)
	}

	localLogger.Infof("index result to es: %+v", experience)

	d.Ack(false)
}
