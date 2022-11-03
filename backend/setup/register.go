package setup

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/bzero/proto/pb_demo"
	"github.com/org/bzero/server"
	"google.golang.org/grpc"
)

func RegisterGrpc(s *grpc.Server) {
	pb_demo.RegisterDemoServiceServer(s, server.NewDemoServer())
}

func RegisterGateway(ctx context.Context, conn *grpc.ClientConn, mux *runtime.ServeMux) {
	pb_demo.RegisterDemoServiceHandler(ctx, mux, conn)
}
