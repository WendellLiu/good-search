package main

import (
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	"github.com/wendellliu/good-search/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file")
	}

	config.Load()
	mongo.Load()
	logger.Load()

	logger.Logger.Info("start")

	companyID := "577f9f9cd4c98d1f8749eecf"
	company := dto.GetCompany(mongo.DB, &companyID)
	logger.Logger.WithFields(logrus.Fields{"company": company}).Info("get result")

	var capital int = 500000
	cursorID, _ := primitive.ObjectIDFromHex("577f9f9cd4c98d1f8749efbd")
	options := dto.Options{
		Limit:    5,
		CursorID: cursorID,
	}

	companies := dto.GetCompanies(mongo.DB, &dto.CompaniesParams{Capital: &capital}, options)
	for i, company := range companies {

		logger.Logger.WithFields(
			logrus.Fields{"index": i},
		).WithFields(
			logrus.Fields{"companies": company},
		).Info("get result")

	}

	experienceID := "598075e1185cc200046fde29"
	experience := dto.GetExperience(mongo.DB, &experienceID)
	logger.Logger.WithFields(logrus.Fields{"experience": experience}).Info("get result")

	_type := "work"
	cursorID, _ = primitive.ObjectIDFromHex("5980a2f5185cc200046fe759")
	options = dto.Options{
		Limit:    5,
		CursorID: cursorID,
	}

	experiences := dto.GetExperiences(mongo.DB, &dto.ExperiencesParams{Type: &_type}, options)
	for i, experience := range experiences {

		logger.Logger.WithFields(
			logrus.Fields{"index": i},
		).WithFields(
			logrus.Fields{"experiences": experience},
		).Info("get result")

	}
}
