package queue

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/logger"
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
	registerUpdateExpConsumer(dependencies, 10, &qu)

	return qu, nil
}

func (q *Queue) NewChannel() (*amqp.Channel, error) {
	channel, err := q.Conn.Channel()

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func registerUpdateExpConsumer(dep *consumer.Dependencies, workerNum int, qu *Queue) error {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "Register UpdateExperienceConsumer"},
	)

	ch, err := qu.NewChannel()

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		UPDATE_EXPERIENCE_QUEUE, // name
		true,                    // durable
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

	for i := 0; i < workerNum; i++ {
		localLogger.Infof("update experience worker: %d", i)
		go func(workerId int, msgCh <-chan amqp.Delivery) {
			for d := range msgCh {
				localLogger.Debugf("update experience by worker: %d", workerId)
				consumer.UpdateExperienceConsumer(dep, &d)
			}

		}(i, msgs)
	}

	return nil
}
