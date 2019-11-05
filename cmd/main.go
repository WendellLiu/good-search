package main

import (
	"os"

	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/mongo/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Load()
	mongo.Load()
	log.Info("start")

	name := "富台機械開發建設有限公司"
	company := dto.GetCompany(&dto.CompanyParams{Name: &name})
	log.WithFields(log.Fields{"company": company}).Info("get result")

	var capital int64 = 500000
	head, _ := primitive.ObjectIDFromHex("577f9f9cd4c98d1f8749eecf")
	options := dto.Options{
		Limit: 10,
		Head:  head,
	}

	companies := dto.GetCompanies(&dto.CompanyParams{Capital: &capital}, options)
	log.WithFields(log.Fields{"companies": companies}).Info("get result")
}
