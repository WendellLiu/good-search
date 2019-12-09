package handlers

import (
	"github.com/wendellliu/good-search/pkg/dto"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

type Server struct {
	pb.UnimplementedGoodSearchServer
	Repository dto.DTO
}
