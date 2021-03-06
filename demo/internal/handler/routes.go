// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	demo1 "github.com/l6l6ng/rabbitmq-go-demo/demo/internal/handler/demo1"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: demo1.DemoHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from2",
				Handler: From2Handler(serverCtx),
			},
		},
	)
}
