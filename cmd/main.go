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

	name := "富台機械開發建設有限公司"
	fmt.Printf("company, %+v \n", dto.GetCompany(&dto.CompanyParams{Name: &name}))

	var capital int64 = 500000
	fmt.Printf("companies, %+v \n", dto.GetCompanies(&dto.CompanyParams{Capital: &capital}, 10))
}
