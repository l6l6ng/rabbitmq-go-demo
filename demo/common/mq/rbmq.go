package mq

import (
	"bytes"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type MqConfig struct {
	Ch *amqp.Channel
}

func newMq(dataSource string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(dataSource)

	failOnError(err, "Failed to connect to RabbitMQ")

	return conn, nil
}

func Connect(dataSource string) *MqConfig {
	Conn, _ := newMq(dataSource)
	ch, err := Conn.Channel()

	failOnError(err, "Failed to open a channel")

	return &MqConfig{
		Ch: ch,
	}

}

func (mq *MqConfig) Publish(name string, body string) bool {
	q, err := mq.Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	err = mq.Ch.Publish(
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

func (mq *MqConfig) PublishWork(name string, body string) bool {
	q, err := mq.Ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	err = mq.Ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	return true
}

func (mq *MqConfig) Consume( name string) {
	q, err := mq.Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

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

func (mq *MqConfig) ConsumeWork( name string) {
	q, err := mq.Ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = mq.Ch.Qos(
		1,
		0,
		false)

	msgs, err := mq.Ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body,[]byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
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
