package es

import (
	"errors"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/wendellliu/good-search/pkg/config"
)

func New() (*elasticsearch.Client, error) {
	esAddress := config.Config.ESAddress
	cfg := elasticsearch.Config{
		Addresses: []string{esAddress},
	}
	es, err := elasticsearch.NewClient(cfg)

	resp, err := es.Info()
	if resp == nil {
		err = errors.New("elasticsearch initial error")
	}

	return es, err
}
