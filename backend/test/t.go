package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"github.com/canbefree/tools/helper"
)

var certFile = "/run/secrets/app_crt"
var keyFile = "/run/secrets/app_key"

func main() {
	l, err := net.Listen("tcp", ":8083")
	helper.PaincErr(err)
	mux := http.DefaultServeMux
	_, err = tls.LoadX509KeyPair(certFile, keyFile)
	helper.PaincErr(err)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})
	http.ServeTLS(l, mux, certFile, keyFile)
}

func log(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
