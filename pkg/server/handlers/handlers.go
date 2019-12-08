package handlers

import (
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

type Server struct {
	pb.UnimplementedGoodSearchServer
	DB dbAdapter.Database
}
