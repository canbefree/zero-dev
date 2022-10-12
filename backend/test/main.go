package main

import (
	"context"
	"net"

	"github.com/canbefree/tools/helper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/repo/proto/pb_demo"
	"github.com/org/repo/server"
	"google.golang.org/grpc"
)

var certFile = "/run/secrets/app_crt"
var keyFile = "/run/secrets/app_key"

func RegisterGrpc(s *grpc.Server) {
	pb_demo.RegisterDemoServiceServer(s, server.NewDemoServer())
}

func RegisterGateway(ctx context.Context, conn *grpc.ClientConn, mux *runtime.ServeMux) {
	pb_demo.RegisterDemoServiceHandler(ctx, mux, conn)
}

func main() {
	l, err := net.Listen("tcp", ":8081")
	helper.PaincErr(err)
	go func() {
		var serverOpts []grpc.ServerOption = GetCred()
		s := grpc.NewServer(serverOpts...)
		defer s.GracefulStop()

		RegisterGrpc(s)
		helper.PaincErr(s.Serve(l))
	}()
}

func makeGrpcServer(address string) (*grpc.ClientConn, *grpc.Server) {
	//
	panic("")
}
