package proxy

import (
	"flag"

	"github.com/elazarl/goproxy"
)

var Defaults = Proxy_config{":8080", false}

func main() {
	var proxy = goproxy.NewProxyHttpServer()
	var verbose = flag.Bool("v", Defaults.Verbose, "activate requess logging to stdout")
	var addr = flag.String("addr", Defaults.Addr, "proxy listen address")

	flag.Parse()
	proxy.Verbose = *verbose

	var proxy_config = Proxy_config{*addr, *verbose}

	Start(proxy_config)
}
