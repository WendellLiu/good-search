package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	"github.com/wendellliu/good-search/pkg/mongo"
	"github.com/wendellliu/good-search/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file")
	}

	config.Load()
	logger.Load()
	server.Server()

	db, err := mongo.New()

	logger.Logger.Info("start")

	companyID := "577f9f9cd4c98d1f8749eecf"
	company, err := dto.GetCompany(context.Background(), db, companyID)

	if err != nil {
		logger.Logger.Error(err)
	}
	logger.Logger.WithFields(logrus.Fields{"company": company}).Info("get result")

	cursorID := "577f9f9cd4c98d1f8749efbd"
	options := dbAdapter.Options{
		Limit:    5,
		CursorID: cursorID,
	}
	var capital int = 500000
	params := &dto.CompaniesParams{Capital: &capital}

	companies, err := dto.GetCompanies(context.Background(), db, params, options)
	if err != nil {
		logger.Logger.Error(err)
	}
	for i, company := range companies {

		logger.Logger.WithFields(
			logrus.Fields{"index": i},
		).WithFields(
			logrus.Fields{"companies": company},
		).Info("get result")

	}

	experienceID := "598075e1185cc200046fde29"
	experience, err := dto.GetExperience(context.Background(), db, experienceID)

	if err != nil {
		logger.Logger.Error(err)
	}
	logger.Logger.WithFields(logrus.Fields{"experience": experience}).Info("get result")

	_type := "work"
	cursorID = "598075e1185cc200046fde29"
	options = dbAdapter.Options{
		Limit:    5,
		CursorID: cursorID,
	}

	experiences, err := dto.GetExperiences(context.Background(), db, &dto.ExperiencesParams{Type: &_type}, options)
	for i, experience := range experiences {

		logger.Logger.WithFields(
			logrus.Fields{"index": i},
		).WithFields(
			logrus.Fields{"experiences": experience},
		).Info("get result")

	}
}
