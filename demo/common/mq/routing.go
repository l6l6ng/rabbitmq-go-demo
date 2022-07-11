package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (mq *MqConfig) PublishRouting(routing_key, body string) bool {
	err := mq.Ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	err = mq.Ch.Publish(
		"logs_direct",
		routing_key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	return true
}

func (mq *MqConfig) ConsumeRouting(routing_key string) {
	err := mq.Ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	q, err := mq.Ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	//for _, routing_key := range []string{"red", "black"} {
		err = mq.Ch.QueueBind(
			q.Name,
			routing_key,
			"logs_direct",
			false,
			nil,
		)
		failOnError(err, "Failed to bind a queue")
	//}

	msgs, err := mq.Ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for message. To exit press Ctrl+C")

	<-forever
}
