package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"github.com/wendellliu/good-search/pkg/queue"
)

func (s *Server) UpdateExperience(ctx context.Context, req *pb.UpdateExperienceReq) (*pb.UpdateExperienceResp, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "UpdateExperience"},
	)

	ch, err := s.Queue.NewChannel()
	defer ch.Close()

	if err != nil {
		localLogger.Error(err)
		return &pb.UpdateExperienceResp{
			Status: pb.Status_FAILURE,
		}, err
	}

	q, err := ch.QueueDeclare(
		queue.UPDATE_EXPERIENCE_QUEUE, // name
		true,                          // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)

	if err != nil {
		localLogger.Error(err)
		return &pb.UpdateExperienceResp{
			Status: pb.Status_FAILURE,
		}, err
	}

	body := req.GetId()
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)

	if err != nil {
		localLogger.Error(err)
		return &pb.UpdateExperienceResp{
			Status: pb.Status_FAILURE,
		}, err
	}

	return &pb.UpdateExperienceResp{
		Status: pb.Status_SUCCESS,
	}, nil
}
