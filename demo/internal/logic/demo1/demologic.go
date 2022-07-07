package demo1

import (
	"context"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/svc"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/types"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	rand.Seed(time.Now().Unix())
	r := 9999 - rand.Intn(8999)
	msg := "hello word!" + strconv.Itoa(r)
	l.svcCtx.Mq.Publish("test-queue1", msg)

	return &types.Response{
		Message: msg,
	}, nil
}
