package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) UpdateExperience(ctx context.Context, req *pb.UpdateExperienceReq) (*pb.UpdateExperienceResp, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "UpdateExperience"},
	)
	experience, err := s.Repository.GetExperience(context.Background(), req.Id)

	err = s.Es.IndexExperience(ctx, experience)

	if err != nil {
		logger.Logger.Error(err)
	}

	localLogger.Infof("index result to es: %+v", experience)
	return &pb.UpdateExperienceResp{
		Status: pb.Status_SUCCESS,
		Experience: &pb.ExperiencePayload{
			Type: experience.Type,
		},
	}, nil
}
