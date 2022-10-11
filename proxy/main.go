package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/canbefree/tools/helper"
	"github.com/elazarl/goproxy"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	// certFile = "../"
	if tls_on() {
		proxy.Tr.ForceAttemptHTTP2 = true
		// 如果使用 tls 需要走 http2协议
		ln, err := net.Listen("tcp", ":3012")
		helper.PaincErr(err)
		// http2.ConfigureServer()
		srv := &http.Server{
			Handler: h2c.NewHandler(proxy, &http2.Server{}),
		}
		http2.ConfigureServer(srv, &http2.Server{})
		err = srv.ServeTLS(ln, certFile, keyFile)
		// err = http.ServeTLS(ln, proxy, certFile, keyFile)
		helper.PaincErr(err)
		return
	}
	http.ListenAndServe(":3010", proxy)
}

func tls_on() bool {
	return true
}
