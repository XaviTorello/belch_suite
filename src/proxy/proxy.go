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

var proxy = goproxy.NewProxyHttpServer()

func Start(params Proxy_config) {

	proxy.Verbose = params.Verbose

	log.Fatal(http.ListenAndServe(params.Addr, proxy))
}
