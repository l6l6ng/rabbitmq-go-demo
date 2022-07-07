package logic

import (
	"context"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type From2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFrom2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *From2Logic {
	return &From2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *From2Logic) From2() error {
	// todo: add your logic here and delete this line
	l.svcCtx.Mq.Consume("test-queue1")
	return nil
}
