package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/logic"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/svc"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/types"
	"net/http"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
