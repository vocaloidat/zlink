package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zlink/zurl/internal/logic"
	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"
)

func ConvertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConvertRequest
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

		l := logic.NewConvertLogic(r.Context(), svcCtx)
		resp, err := l.Convert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
