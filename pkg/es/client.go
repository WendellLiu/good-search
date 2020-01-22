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

	res, err := req.Do(ctx, es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return err
}

func (es *Elasticsearch) SearchExperiences(ctx context.Context, keyword string) (
	experienceIds []string,
	err error,
) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "es-SearchExperiences"},
	)
	client := es.Client

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": config.Config.Search.ExperiencesSearch.Fields,
			},
		},
		"_source": []string{"title"},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		localLogger.Fatalf("Error encoding query: %s", err)
	}

	req := esapi.SearchRequest{
		Index:          []string{EXPERIENCE_INDEX},
		Body:           &buf,
		TrackTotalHits: true,
	}

	res, err := req.Do(ctx, client)

	localLogger.Infof("res: %+v \n", res)
	return []string{"123", "321"}, nil
}
