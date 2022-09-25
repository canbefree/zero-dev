package main

import (
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		fmt.Printf("request:%v", r.Body)
		return r, nil
	})
	// var certFile, keyFile string
	// certFile = "../"
	http.ListenAndServe(":3010", proxy)
}
