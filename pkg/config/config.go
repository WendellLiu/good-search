package config

import (
	"os"
)

type BasicConfig struct {
	MongoDBHost     string
	MongoDBPort     string
	MongoDBPassword string
	MongoDBName     string
	GRPCPort        string
	ESAddress       string
}

var Config BasicConfig

func Load() {
	Config = BasicConfig{
		MongoDBHost:     os.Getenv("MONGO_DB_HOST"),
		MongoDBPort:     os.Getenv("MONGO_DB_PORT"),
		MongoDBPassword: os.Getenv("MONGO_DB_PASSWORD"),
		MongoDBName:     os.Getenv("MONGO_DB_NAME"),
		GRPCPort:        os.Getenv("GRPC_PORT"),
		ESAddress:       os.Getenv("ES_ADDRESS"),
	}
}
