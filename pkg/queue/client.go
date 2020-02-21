package queue

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/config"
)

const (
	UPDATE_EXPERIENCE_QUEUE = "UPDATE_EXPERIENCE_QUEUE"
)

type Queue struct {
	Conn *amqp.Connection
}

func New() (Queue, error) {
	amqpURI := fmt.Sprintf("amqp://guest:guest@localhost:%s/", config.Config.Rabbitmq.Port)
	conn, err := amqp.Dial(amqpURI)

	if err != nil {
		panic(err)
	}

	return Queue{
		Conn: conn,
	}, nil
}
