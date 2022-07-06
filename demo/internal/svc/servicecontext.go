package svc

import (
	"github.com/l6l6ng/rabbitmq-go-demo/demo/common/mq"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ServiceContext struct {
	Config config.Config
	Mq     *amqp.Channel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, _ := amqp.Dial(c.RabbitMq.DataSource)
	defer conn.Close()

	return &ServiceContext{
		Config: c,
		Mq:     mq.Connect(c.RabbitMq.DataSource),
	}
}
