package wechat

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

/**
实现Response的响应对象
*/
type wechatResponseWrapper struct {
	*fasthttp.RequestCtx
}

func (w wechatResponseWrapper) Header() http.Header {
	return map[string][]string{}
}

func (w wechatResponseWrapper) Write(bytes []byte) (int, error) {
	return w.RequestCtx.Write(bytes)
}

func (w wechatResponseWrapper) WriteHeader(statusCode int) {
	w.Response.SetStatusCode(statusCode)
}
