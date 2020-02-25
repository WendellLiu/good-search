package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config SystemConfig

type SystemConfig struct {
	DevelopmentMode bool           `yaml:"development_mode"`
	Search          SearchConfig   `yaml:"search"`
	Mongo           MongoConfig    `yaml:"mongo"`
	Es              EsConfig       `yaml:"es"`
	Grpc            GrpcConfig     `yaml:"grpc"`
	Rabbitmq        RabbitmqConfig `yaml:"rabbitmq"`
}

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type EsConfig struct {
	Address string `yaml:"address"`
}

type GrpcConfig struct {
	Port string `yaml:"port"`
}

type RabbitmqConfig struct {
	Port string `yaml:"port"`
}

type ExperiencesSearch struct {
	Fields   []string `yaml:"fields"`
	Analyzer string   `yaml:"analyzer"`
	Type     string   `yaml:"type"`
}

type SearchConfig struct {
	ExperiencesSearch ExperiencesSearch `yaml:"experiences_search"`
}

func readConfigYaml() (systemConfig SystemConfig, err error) {
	c := SystemConfig{}
	bytes, err := ioutil.ReadFile("config.yml")

	if err != nil {
		return c, err
	}

	yaml.Unmarshal(bytes, &c)

	return c, nil
}

func Load() error {
	systemConfig, err := readConfigYaml()

	if err != nil {
		return err
	}

	Config = systemConfig
	return nil
}
