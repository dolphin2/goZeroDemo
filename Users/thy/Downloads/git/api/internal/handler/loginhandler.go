package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroDemo/Users/thy/Downloads/git/api/internal/logic"
	"goZeroDemo/Users/thy/Downloads/git/api/internal/svc"
	"goZeroDemo/Users/thy/Downloads/git/api/internal/types"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
