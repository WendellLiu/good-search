package handlers

import (
	"context"

	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) UpdateExperience(ctx context.Context, req *pb.UpdateExperienceReq) (*pb.UpdateExperienceResp, error) {

	//if err != nil {
	//logger.Logger.Error(err)
	//}
	return &pb.UpdateExperienceResp{Status: pb.Status_SUCCESS}, nil
}
