package handlers

import (
	"context"

	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) SearchExperiences(ctx context.Context, req *pb.SearchExperiencesReq) (*pb.SearchExperiencesResp, error) {
	experienceIds, err := s.Es.SearchExperiences(ctx, req.Keyword)

	if err != nil {
		return &pb.SearchExperiencesResp{
			Status: pb.Status_FAILURE,
		}, err
	}

	return &pb.SearchExperiencesResp{
		Status:        pb.Status_SUCCESS,
		ExperienceIds: experienceIds,
	}, nil
}
