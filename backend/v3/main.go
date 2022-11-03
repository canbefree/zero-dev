package main

import (
	"context"
	"net"

	"github.com/canbefree/tools/helper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/bzero/setup"
	"github.com/org/bzero/vars"
	"google.golang.org/grpc"
)

func main() {
	var grpcServer GRPCServerIFace
	var gatewayServer GatewayServerIFace
	if vars.GRPC_TLS_ON {
		// grpcServer =
	} else {

	}
	if vars.HTTP_TLS_ON {

	} else {

	}

	server := NewServer(grpcServer, gatewayServer)
	go func() {
		helper.PaincErr(server.StartServer())
	}()
	helper.PaincErr(server.Listen())
}

type Server struct {
	grpcEndpoint    string
	gatewayEndPoint string

	serverOptions []grpc.ServerOption
	clientOptions []grpc.DialOption

	GRPCServerIFace
	GatewayServerIFace
}

type GRPCServerIFace interface {
	SetEndPoint()
	SetServerOpts()
	StartServer() error
}

type GRPCServer struct {
}

type GatewayServerIFace interface {
	SetClientDialOpts()
	SetEndPoint()
	Listen() error
}

func (s *Server) StartServer() error {
	l, err := net.Listen("tcp", s.grpcEndpoint)
	helper.PaincErr(err)
	gsever := grpc.NewServer(s.serverOptions...)
	defer gsever.GracefulStop()
	setup.RegisterGrpc(gsever)
	return gsever.Serve(l)
}

func (s *Server) Listen() error {
	mux := runtime.NewServeMux()
	ctx := context.Background()

	conn, err := grpc.Dial(s.grpcEndpoint, s.clientOptions...)
	if err != nil {
		return err
	}
	setup.RegisterGateway(ctx, conn, mux)
	return s.Listen()
}

func NewServer(grpcServer GRPCServerIFace, gatewayServer GatewayServerIFace) *Server {
	return &Server{
		GRPCServerIFace:    grpcServer,
		GatewayServerIFace: gatewayServer,
	}
}
