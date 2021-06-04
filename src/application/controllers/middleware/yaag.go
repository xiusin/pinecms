package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/betacraft/yaag/yaag"
	"github.com/betacraft/yaag/yaag/models"
	"github.com/valyala/fasthttp"
	"github.com/xiusin/pine"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/url"
	"strings"
)

func YaagApiDoc() pine.Handler {
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "PineCMS ApiDoc",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	return func(ctx *pine.Context) {
		if !yaag.IsOn() {
			ctx.Next()
			return
		}
		apiCall := models.ApiCall{}
		requestBefore(&apiCall, ctx.RequestCtx)
		ctx.Next()

		if yaag.IsStatusCodeValid(ctx.Response.StatusCode()) {
			apiCall.MethodType = string(ctx.Method())
			apiCall.CurrentPath = strings.Split(string(ctx.RequestURI()), "?")[0]
			apiCall.ResponseBody = string(ctx.Response.Body())
			apiCall.ResponseCode = ctx.Response.StatusCode()
			headers := map[string]string{}
			//for k, v := range  {
			//	log.Println(k, v)
			//	headers[k] = strings.Join(v, " ")
			//}
			apiCall.ResponseHeader = headers
			go yaag.GenerateHtml(&apiCall)
		}
	}
}

//func handleMultipart(apiCall *models.ApiCall, req *http.Request) {
//	apiCall.RequestHeader["Content-Type"] = "multipart/form-data"
//	req.ParseMultipartForm(MaxInMemoryMultipartSize)
//	apiCall.PostForm = ReadMultiPostForm(req.MultipartForm)
//}

func readHeaders(req *fasthttp.RequestCtx) map[string]string {
	var m = map[string]string{}
	err := json.Unmarshal(req.Request.Header.Header(), &m)
	if err != nil {
		pine.Logger().Error("解析header错误", string(req.Request.Header.Header()))
	}
	return m
}

func readQueryParams(req *fasthttp.RequestCtx) map[string]string {
	params := map[string]string{}
	u, err := url.Parse(string(req.RequestURI()))
	if err != nil {
		return params
	}
	for k, v := range u.Query() {
		if len(v) < 1 {
			continue
		}
		params[k] = v[0]
	}
	return params
}

func readMultiPostForm(mpForm *multipart.Form) map[string]string {
	postForm := map[string]string{}
	for key, val := range mpForm.Value {
		postForm[key] = val[0]
	}
	return postForm
}

// One of the copies, say from b to r2, could be avoided by using a more
// elaborate trick where the other copy is made during Request/Response.Write.
// This would complicate things too much, given that these functions are for
// debugging only.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, nil, err
	}
	if err = b.Close(); err != nil {
		return nil, nil, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func readPostForm(req *fasthttp.RequestCtx) map[string]string {
	postForm := map[string]string{}
	for _, param := range strings.Split(string(req.PostBody()), "&") {
		value := strings.Split(param, "=")
		postForm[value[0]] = value[1]
	}
	return postForm
}

func requestBefore(apiCall *models.ApiCall, req *fasthttp.RequestCtx) {
	apiCall.RequestHeader = readHeaders(req)
	apiCall.RequestUrlParams = readQueryParams(req)
	val, ok := apiCall.RequestHeader["Content-Type"]
	log.Println(val)
	if ok {
		ct := strings.TrimSpace(apiCall.RequestHeader["Content-Type"])
		switch ct {
		case "application/x-www-form-urlencoded":
			fallthrough
		case "application/json, application/x-www-form-urlencoded":
			log.Println("Reading form")
			apiCall.PostForm = readPostForm(req)
		case "application/json":
			log.Println("Reading body")
			apiCall.RequestBody = string(req.PostBody())
		default:
			apiCall.RequestBody = string(req.PostBody())
		}
	}
}
