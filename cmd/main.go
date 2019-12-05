package main

import (
	"github.com/wendellliu/good-search/pkg/config"
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

	db, err := mongo.New()
	server.Server(db)

	//experienceID := "598075e1185cc200046fde29"
	//experience, err := dto.GetExperience(context.Background(), db, experienceID)

	//if err != nil {
	//logger.Logger.Error(err)
	//}
	//logger.Logger.WithFields(logrus.Fields{"experience": fmt.Sprintf("%+v", experience)}).Info("get result")

	//_type := "work"
	//cursorID = "598075e1185cc200046fde29"
	//options = dbAdapter.Options{
	//Limit:    5,
	//CursorID: cursorID,
	//}

	//experiences, err := dto.GetExperiences(context.Background(), db, &dto.ExperiencesParams{Type: &_type}, options)
	//for i, experience := range experiences {

	//logger.Logger.WithFields(
	//logrus.Fields{"index": i},
	//).WithFields(
	//logrus.Fields{"experiences": experience},
	//).Info("get result")

	//}
}
