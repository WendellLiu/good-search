package config

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/logger"

	"gopkg.in/yaml.v3"
)

var Config BasicConfig

type BasicConfig struct {
	SystemConfig
	MongoDBHost     string
	MongoDBPort     string
	MongoDBPassword string
	MongoDBName     string
	GRPCPort        string
	ESAddress       string
}

type SystemConfig struct {
	Search SearchConfig `yaml:"search"`
}

type ExperiencesSearch struct {
	Fields   []string `yaml:"fields"`
	Analyzer string   `yaml:"analyzer"`
	Type     string   `yaml:"type"`
}

type SearchConfig struct {
	ExperiencesSearch ExperiencesSearch `yaml:"experiences_search"`
}

func readConfigYaml() SystemConfig {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "config-readConfigYaml"},
	)

	c := SystemConfig{}
	bytes, err := ioutil.ReadFile("config.yml")

	if err != nil {
		localLogger.Error("read system config error")
		return c
	}

	yaml.Unmarshal(bytes, &c)

	return c
}

func Load() {
	systemConfig := readConfigYaml()
	Config = BasicConfig{
		SystemConfig:    systemConfig,
		MongoDBHost:     os.Getenv("MONGO_DB_HOST"),
		MongoDBPort:     os.Getenv("MONGO_DB_PORT"),
		MongoDBPassword: os.Getenv("MONGO_DB_PASSWORD"),
		MongoDBName:     os.Getenv("MONGO_DB_NAME"),
		GRPCPort:        os.Getenv("GRPC_PORT"),
		ESAddress:       os.Getenv("ES_ADDRESS"),
	}
}
