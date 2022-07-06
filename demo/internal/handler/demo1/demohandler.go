package demo1

import (
	"net/http"

	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/logic/demo1"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/svc"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := demo1.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
