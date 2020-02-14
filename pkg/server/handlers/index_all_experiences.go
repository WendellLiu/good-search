package handlers

import (
	"context"
	"math"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) indexToES(ctx context.Context, experiences []dto.Experience) (err error) {
	for _, experience := range experiences {
		err = s.Es.IndexExperience(ctx, experience)
		if err != nil {
			break
		}
	}
	return err
}

func (s *Server) IndexAllExperiences(ctx context.Context, req *pb.IndexAllExperiencesReq) (*pb.IndexAllExperiencesResp, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "IndexAllExperiences"},
	)

	workerNum := 10
	offset := int64(5)
	count, err := s.Repository.GetExperiencesCount(ctx)
	if err != nil {
		localLogger.Error(err)
		return &pb.IndexAllExperiencesResp{
			Status: pb.Status_FAILURE,
		}, nil
	}
	var nilObjectID primitive.ObjectID
	firstID := nilObjectID.Hex()

	bufferCount := int(math.Ceil(float64(count) / float64(offset)))
	lookupIds := make(chan string, bufferCount)

	var wg sync.WaitGroup
	var mux sync.Mutex

	counter := 0

	// for dev
	devCounter := 2000

	lookupIds <- firstID
	for i := 0; i < workerNum; i++ {
		localLogger.Infof("worker: %d", i)
		wg.Add(1)
		go func(candidateIds chan string) {
			for id := range candidateIds {
				localLogger.Infof("id: %s", id)
				// for dev
				mux.Lock()
				devCounter--
				mux.Unlock()

				experiences, err := s.Repository.GetExperiences(
					ctx,
					&dto.ExperiencesParams{},
					dbAdapter.Options{
						Limit:    offset,
						CursorID: id,
					},
				)
				if err != nil {
					localLogger.Error(err)
					break
				}

				lenExps := len(experiences)

				if devCounter <= 0 || lenExps < 1 {
					close(candidateIds)
					localLogger.Infof("over and return")
					break
				}

				lastExp := experiences[lenExps-1]
				lastCursor := lastExp.ID.Hex()

				candidateIds <- lastCursor
				err = s.indexToES(ctx, experiences)

				if err != nil {
					localLogger.Error(err)
				}
				mux.Lock()
				counter = counter + len(experiences)
				mux.Unlock()

				localLogger.Infof("index experiences start with : %s(offset: %d)", id, offset)
			}
			wg.Done()
		}(lookupIds)
	}

	wg.Wait()

	localLogger.Infof("index %d documents", counter)

	return &pb.IndexAllExperiencesResp{
		Status: pb.Status_SUCCESS,
	}, nil
}
