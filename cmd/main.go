package main

import (
	"fmt"
	"log"

	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/mongo/dto"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Load()
	mongo.Load()
	fmt.Println("start")
	fmt.Println(dto.GetCompany())
}
