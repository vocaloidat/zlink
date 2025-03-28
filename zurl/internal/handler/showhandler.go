package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zlink/zurl/bodyx"
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

		// 自定义参数规则校验
		err := validator.New().StructCtx(r.Context(), &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("长连接参数错误"))
			return
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			// 自定义结构返回
			bodyx.Writer(w, r, resp, err)
		} else {
			// StatusMovedPermanently
			// 跳转 301永久重定向，302临时重定向
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
			//http.Redirect(w, r, resp.LongUrl, http.StatusMovedPermanently)
		}

	}
}
