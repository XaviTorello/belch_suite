package proxy

import (
	"flag"
)

var Defaults = Proxy_config{":8080", false}

func main() {
	var verbose = flag.Bool("v", Defaults.Verbose, "activate requess logging to stdout")
	var addr = flag.String("addr", Defaults.Addr, "proxy listen address")

	flag.Parse()

	var proxy_config = Proxy_config{*addr, *verbose}

	Start_simple(proxy_config)
	// Start_interceptor(proxy_config)
}
