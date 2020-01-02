package handlers

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) UpdateExperience(ctx context.Context, req *pb.UpdateExperienceReq) (*pb.UpdateExperienceResp, error) {
	experience, err := s.Repository.GetExperience(context.Background(), req.Id)

	err = s.Es.IndexExperience(context.Background(), experience)

	if err != nil {
		logger.Logger.Error(err)
	}

	logger.Logger.WithFields(logrus.Fields{"experience": fmt.Sprintf("%+v", experience)}).Info("get result")
	return &pb.UpdateExperienceResp{
		Status: pb.Status_SUCCESS,
		Experience: &pb.ExperiencePayload{
			Type: experience.Type,
		},
	}, nil
}
