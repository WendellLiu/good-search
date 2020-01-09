package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) IndexAllExperiences(ctx context.Context, req *pb.IndexAllExperiencesReq) (*pb.IndexAllExperiencesResp, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "IndexAllExperiences"},
	)

	go func() {
		offset := int64(2)
		limitCounter := int64(10)
		var nilObjectID primitive.ObjectID
		currentCursor := nilObjectID.Hex()

		for limitCounter > 0 {
			experiences, err := s.Repository.GetExperiences(
				ctx,
				&dto.ExperiencesParams{},
				dbAdapter.Options{
					Limit:    offset,
					CursorID: currentCursor,
				},
			)

			if err != nil {
				localLogger.Error(err)

				break
			}

			for _, experience := range experiences {
				err = s.Es.IndexExperience(ctx, experience)
				if err != nil {
					localLogger.Error(err)
					break
				}
				localLogger.Infof("index experience id: %+v", experience.ID)
				currentCursor = experience.ID.Hex()
			}

			if err != nil {
				localLogger.Error(err)

				break
			}

			//post-process
			limitCounter = limitCounter - offset
		}
	}()

	return &pb.IndexAllExperiencesResp{
		Status: pb.Status_SUCCESS,
	}, nil
}
