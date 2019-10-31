package config

import "os"

type BasicConfig struct {
	foo string
}

var Config BasicConfig

func LoadConfig() {
	Config = BasicConfig{
		foo: os.Getenv("FOO"),
	}
}
