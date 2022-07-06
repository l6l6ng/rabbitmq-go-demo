package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type MqConfig struct {
}

var Conn *amqp.Connection
var Chan *amqp.Channel

func newMq(dataSource string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(dataSource)

	failOnError(err, "Failed to connect to RabbitMQ")

	return conn, nil
}

func Connect(dataSource string) *amqp.Channel {
	Conn, _ := newMq(dataSource)
	Chan, err := Conn.Channel()

	failOnError(err, "Failed to open a channel")

	return Chan

}

func Publish(chan1 *amqp.Channel, name string, body string) bool {
	q, err := chan1.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	err = chan1.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	return true
}

func Consume(chan1 *amqp.Channel, name string) {
	q, err := chan1.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := chan1.Consume(
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
