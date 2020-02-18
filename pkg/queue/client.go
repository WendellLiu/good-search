package queue

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/config"
)

type Queue struct {
	Conn *amqp.Connection
}

func New() (Queue, error) {
	amqpURI := fmt.Sprintf("amqp://guest:guest@localhost:%s/", config.Config.Rabbitmq.Port)
	fmt.Printf("ampqURI: %s \n", amqpURI)
	conn, err := amqp.Dial(amqpURI)

	if err != nil {
		return Queue{}, err
	}

	return Queue{
		Conn: conn,
	}, nil
}
