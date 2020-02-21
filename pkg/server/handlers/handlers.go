package handlers

import (
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"github.com/wendellliu/good-search/pkg/queue"
)

type Server struct {
	pb.UnimplementedGoodSearchServer
	Repository dto.DTO
	Es         es.Elasticsearch
	Queue      queue.Queue
}
