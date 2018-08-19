package main

import (
	"fmt"
	"os"
	"github.com/elazarl/goproxy"
    "log"
	// "path/filepath"
)

func openFile(path string){
	os.Open(path)
}

type Request struct {
	IP string
	Port int
	Uri string
	Payload string
}

func start_proxy() {
	proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
	return proxy
}

func main() {
	Requests_list := map[int]Request{}
	// Requests_list := []Request{}

	Requests_list[0] =  Request{"8.8.8.8", 80, "/", "payload"}
	Requests_list[-1] =  Request{"4.4.4.4", 80, "/", "payload"}

	for _, request := range Requests_list {
		fmt.Println(request.IP)
	}
}

