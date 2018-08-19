package main

import (
	"fmt"
	"os"

	// "github.com/XaviTorello/belch"
	// "path/filepath"
	"proxy"
)

func openFile(path string) {
	os.Open(path)
}

type Request struct {
	IP      string
	Port    int
	Uri     string
	Payload string
}

var proxy_config = proxy.Defaults

var Requests_list = map[int]Request{}

func main() {

	proxy_config.Verbose = true
	proxy.Start(proxy_config)

	// Requests_list := []Request{}
	Requests_list[0] = Request{"8.8.8.8", 80, "/", "payload"}
	Requests_list[-1] = Request{"4.4.4.4", 80, "/", "payload"}

	for _, request := range Requests_list {
		fmt.Println(request.IP)
	}
}
