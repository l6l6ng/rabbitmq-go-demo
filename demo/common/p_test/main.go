package main

import (
	"fmt"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/common/mq"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var dataSource = "amqp://admin:123456@localhost:5672/"
	mqcf := mq.Connect(dataSource)

	i := 0
	msg := ""
	for i < 10 {
		rand.Seed(time.Now().UnixMicro())
		r := 9999 - rand.Intn(8999)
		msg = "hello word.!" + strconv.Itoa(r)
		mqcf.PublishPublish(msg)
		i++
	}
	fmt.Println(i)
}
