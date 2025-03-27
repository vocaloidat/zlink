package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zlink/zurl/internal/logic"
	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"
)

func ZurlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZurlLogic(r.Context(), svcCtx)
		resp, err := l.Zurl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)

		} else {
			//// 返回重定向http响应 方法1
			//// 设置重定向地址
			//w.Header().Set("Location", resp.LongUrl)
			//// 设置重定向状态码
			//w.WriteHeader(http.StatusFound)

			// 返回重定向http响应 第二种方法
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
