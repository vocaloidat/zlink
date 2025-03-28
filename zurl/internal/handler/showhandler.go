package handler

import (
	"net/http"
	"zlink/zurl/bodyx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zlink/zurl/internal/logic"
	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		// 自定义结构返回
		bodyx.Writer(w, r, resp, err)
	}
}
