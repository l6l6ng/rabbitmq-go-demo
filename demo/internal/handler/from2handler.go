package handler

import (
	"net/http"

	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/logic"
	"github.com/l6l6ng/rabbitmq-go-demo/demo/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func From2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFrom2Logic(r.Context(), svcCtx)
		err := l.From2()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
