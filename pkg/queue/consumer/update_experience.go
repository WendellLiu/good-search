package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
)

func UpdateExperienceConsumer(d *amqp.Delivery) {
	fmt.Printf("Received a message: %s \n", d.Body)
	d.Ack(false)
}
