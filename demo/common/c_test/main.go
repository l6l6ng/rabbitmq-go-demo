package main

import (
	"github.com/l6l6ng/rabbitmq-go-demo/demo/common/mq"
	"os"
)

func main() {
	var dataSource = "amqp://admin:123456@localhost:5672/"
	mqcf := mq.Connect(dataSource)

	//发布订阅模式
	//mqcf.ConsumePublish()

	//routing 模式
	routing_key := os.Args[1]
	mqcf.ConsumeRouting(routing_key)
}
