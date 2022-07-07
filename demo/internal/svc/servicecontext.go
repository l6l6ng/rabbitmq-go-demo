package svc

import (
	"github.com/l6l6ng/rabbitmq-go-demo/demo/common/mq"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mq     *mq.MqConfig
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mq:     mq.Connect(c.RabbitMq.DataSource),
	}
}
