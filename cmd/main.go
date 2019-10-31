package main

import (
	"fmt"
	"log"

	"github.com/wendellliu/good-search/pkg/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.LoadConfig()
	fmt.Println("start")

}
