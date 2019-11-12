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

	name := "富台機械開發建設有限公司"
	company := dto.GetCompany(mongo.DB, &dto.CompanyParams{Name: &name})
	logger.Logger.WithFields(logrus.Fields{"company": company}).Info("get result")

	var capital int = 500000
	cursorID, _ := primitive.ObjectIDFromHex("577f9f9cd4c98d1f8749efbd")
	options := dto.Options{
		Limit:    5,
		CursorID: cursorID,
	}

	companies := dto.GetCompanies(mongo.DB, &dto.CompanyParams{Capital: &capital}, options)
	for i, company := range companies {

		logger.Logger.WithFields(
			logrus.Fields{"index": i},
		).WithFields(
			logrus.Fields{"companies": company},
		).Info("get result")

	}
	//var capital int = 500000
	//cursorID, _ := primitive.ObjectIDFromHex("577f9f9cd4c98d1f8749efbd")
	//options := dto.Options{
	//Limit:    5,
	//CursorID: cursorID,
	//}

	//companies := dto.GetCompanies(mongo.DB, &dto.CompanyParams{Capital: &capital}, options)
	//for i, company := range companies {

	//logger.Logger.WithFields(
	//logrus.Fields{"index": i},
	//).WithFields(
	//logrus.Fields{"companies": company},
	//).Info("get result")

	//}
}
