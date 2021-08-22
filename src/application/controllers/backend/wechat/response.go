package wechat

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

type wechatResponseWriter struct {
	*fasthttp.RequestCtx
}

func (w wechatResponseWriter) Header() http.Header {
	return map[string][]string{}
}

func (w wechatResponseWriter) Write(bytes []byte) (int, error) {
	return w.RequestCtx.Write(bytes)
}

func (w wechatResponseWriter) WriteHeader(statusCode int) {
	w.Response.SetStatusCode(statusCode)
}
