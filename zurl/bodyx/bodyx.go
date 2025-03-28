package bodyx

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Bodyx struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Writer(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	body := Bodyx{}
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Data = resp
	}
	httpx.OkJsonCtx(r.Context(), w, body)
}
