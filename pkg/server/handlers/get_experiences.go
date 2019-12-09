package handlers

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/logger"

	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) GetExperience(ctx context.Context, req *pb.GetExperienceReq) (*pb.GetExperienceResp, error) {
	experience, err := s.Repository.GetExperience(context.Background(), req.Id)

	if err != nil {
		logger.Logger.Error(err)
	}
	logger.Logger.WithFields(logrus.Fields{"experience": fmt.Sprintf("%+v", experience)}).Info("get result")
	return &pb.GetExperienceResp{Type: experience.Type}, nil
}
