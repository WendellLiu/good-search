package queue

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/queue/consumer"
)

const (
	UPDATE_EXPERIENCE_QUEUE = "UPDATE_EXPERIENCE_QUEUE"
)

type Queue struct {
	Conn *amqp.Connection
}

func New(dependencies *consumer.Dependencies) (Queue, error) {
	amqpURI := fmt.Sprintf("amqp://guest:guest@localhost:%s/", config.Config.Rabbitmq.Port)
	conn, err := amqp.Dial(amqpURI)

	if err != nil {
		panic(err)
	}

	qu := Queue{
		Conn: conn,
	}

	// register consumers
	registerUpdateExpConsumer(dependencies, &qu)

	return qu, nil
}

func (q *Queue) NewChannel() (*amqp.Channel, error) {
	channel, err := q.Conn.Channel()

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func registerUpdateExpConsumer(dep *consumer.Dependencies, qu *Queue) error {
	ch, err := qu.NewChannel()

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		UPDATE_EXPERIENCE_QUEUE, // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args

	)

	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			consumer.UpdateExperienceConsumer(dep, &d)
		}

	}()

	return nil
}
