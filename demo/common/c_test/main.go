package main

import "github.com/l6l6ng/rabbitmq-go-demo/demo/common/mq"

func main() {
	var dataSource = "amqp://admin:123456@localhost:5672/"
	mqcf := mq.Connect(dataSource)

	mqcf.ConsumePublish()
}
