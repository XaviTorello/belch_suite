package proxy

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

type Proxy_config struct {
	Addr    string
	Verbose bool
}

func Start_simple(params Proxy_config) {

	proxy := goproxy.NewProxyHttpServer()

	log.Println("Starting simple proxy")
	proxy.Verbose = params.Verbose

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("X-BelchSuite", "WTF!r0lf")
			log.Println(r)
			return r, nil
		})

	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			resp.Header.Set("X-BelchSuite", "WTF!r0lf")
			return resp
		})

	log.Fatal(http.ListenAndServe(params.Addr, proxy))
}
