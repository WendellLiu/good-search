package es

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
)

type Elasticsearch struct {
	Client *elasticsearch.Client
}

func New() (Elasticsearch, error) {
	esAddress := config.Config.ESAddress
	cfg := elasticsearch.Config{
		Addresses: []string{esAddress},
	}
	esClient, err := elasticsearch.NewClient(cfg)

	resp, err := esClient.Info()
	if resp == nil {
		err = errors.New("elasticsearch initial error")
	}

	return Elasticsearch{Client: esClient}, err
}

const EXPERIENCE_INDEX = "goodjob-experience"

func (es *Elasticsearch) IndexExperience(ctx context.Context, experience dto.Experience) error {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "es-IndexExperience"},
	)
	var err error

	id := experience.ID.Hex()

	b, err := json.Marshal(experience)
	if err != nil {
		localLogger.Error(err)
	}
	req := esapi.IndexRequest{
		Index:      EXPERIENCE_INDEX,
		DocumentID: id,
		Body:       bytes.NewReader(b),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return err
}
