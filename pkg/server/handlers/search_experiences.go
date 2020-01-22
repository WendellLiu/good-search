package handlers

import (
	"context"

	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) SearchExperiences(ctx context.Context, req *pb.SearchExperiencesReq) (*pb.SearchExperiencesResp, error) {
	s.Es.SearchExperiences(ctx, req.Keyword)
	return &pb.SearchExperiencesResp{
		Status: pb.Status_SUCCESS,
		ExperienceIds: []string{
			"123",
			"321",
		},
	}, nil
}
