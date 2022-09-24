package main

import (
	"context"
	"net"
	"net/http"

	"github.com/canbefree/tools/helper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/repo/proto/pb_demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	s := grpc.NewServer()
	defer s.GracefulStop()

	RegisterGrpc(s)
	l, err := net.Listen("tcp", ":8081")
	helper.PaincErr(err)
	go func() {
		helper.PaincErr(s.Serve(l))
	}()

	mux := runtime.NewServeMux()
	RegisterGateway(mux)
	http.ListenAndServe(":8082", mux)
}

func RegisterGrpc(s *grpc.Server) {
	pb_demo.RegisterDemoServiceServer(s, pb_demo.UnimplementedDemoServiceServer{})
}

func RegisterGateway(mux *runtime.ServeMux) {
	ctx := context.TODO()
	var dailOptions []grpc.DialOption = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	helper.PaincErr(pb_demo.RegisterDemoServiceHandlerFromEndpoint(ctx, mux, ":8081", dailOptions))
}
