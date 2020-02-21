package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"github.com/wendellliu/good-search/pkg/queue"
)

func (s *Server) Dumb(ctx context.Context, req *pb.DumbReq) (*pb.DumbResp, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "Dumb"},
	)

	ch, err := s.Queue.NewChannel()
	defer ch.Close()

	if err != nil {
		localLogger.Error(err)
	}

	q, err := ch.QueueDeclare(
		queue.UPDATE_EXPERIENCE_QUEUE, // name
		false,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	return &pb.DumbResp{}, nil
}
