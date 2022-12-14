package main

import (
	"context"
	"net"
	"net/http"

	"github.com/canbefree/tools/helper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/repo/proto/pb_demo"
	"github.com/org/repo/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var certFile = "/run/secrets/app_crt"
	var keyFile = "/run/secrets/app_key"

	var serverOpts []grpc.ServerOption
	if tls_on() {
		c, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		helper.PaincErr(err)
		serverOpts = append(serverOpts, grpc.Creds(c))
	}

	s := grpc.NewServer(serverOpts...)
	defer s.GracefulStop()

	RegisterGrpc(s)
	l, err := net.Listen("tcp", ":8081")
	helper.PaincErr(err)
	go func() {
		helper.PaincErr(s.Serve(l))
	}()

	mux := runtime.NewServeMux()
	RegisterGateway(mux)

	if tls_on() {
		// 如果使用 tls 需要走 http2协议
		err = http.Serve(l, h2c.NewHandler(
			mux,
			&http2.Server{}),
		)
		// helper.PaincErr(http.ListenAndServeTLS(":8082", certFile, keyFile, mux))
		return
	}
	helper.PaincErr(http.ListenAndServe(":8082", mux))
}

func RegisterGrpc(s *grpc.Server) {
	pb_demo.RegisterDemoServiceServer(s, server.NewDemoServer())
}

func RegisterGateway(mux *runtime.ServeMux) {
	ctx := context.TODO()
	var dailOptions []grpc.DialOption = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	helper.PaincErr(pb_demo.RegisterDemoServiceHandlerFromEndpoint(ctx, mux, ":8081", dailOptions))
}

func tls_on() bool {
	return true
}
