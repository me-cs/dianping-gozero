package apix

import (
	"net/http"

	"github.com/homelight/json"
	"github.com/me-cs/dianping-gozero/library/ecode"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Success bool        `json:"success"`
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJson writes v as json string into w with code.
func writeJson(w http.ResponseWriter, code int, v interface{}) {
	bs, err := json.MarshalSafeCollections(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(httpx.ContentType, "application/json; charset=utf-8")
	w.WriteHeader(code)
	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}

func Response(w http.ResponseWriter, rsp interface{}, err error) {
	var body Body
	if err != nil {
		e := ecode.Error(err)
		body.Success = false
		body.Code = e.Code()
		body.Message = e.Message()
	} else {
		body.Success = true
		body.Message = "OK"
	}
	if rsp != nil {
		body.Data = rsp
	}
	writeJson(w, http.StatusOK, body)
}
