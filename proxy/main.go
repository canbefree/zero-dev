package main

import (
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	var certFile = "/run/secrets/app_crt"
	var keyFile = "/run/secrets/app_key"


	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		fmt.Printf("request:%v", r.Body)
		return r, nil
	})
	// var certFile, keyFile string
	// certFile = "../"
	if tls_on() {
		http.ListenAndServeTLS(":3010", certFile, keyFile, proxy)
		return
	}
	http.ListenAndServe(":3010", proxy)
}

func tls_on() bool {
	return true
}
