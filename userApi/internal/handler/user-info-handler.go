package handler

import (
	"net/http"

	"subTree/userApi/internal/logic"
	"subTree/userApi/internal/svc"
	"subTree/userApi/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.Result{Code: -1, Message: err.Error()})
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(req)
		if err != nil {
			httpx.OkJson(w, err)
		} else {
			httpx.OkJson(w, types.Result{Data: resp})
		}
	}
}
