package main

import (
	"context"
	"net"
	"net/http"

	"github.com/canbefree/tools/helper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/org/bzero/setup"
	"github.com/org/bzero/utils"
	"github.com/org/bzero/vars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	serverOpts := GetServerOpts()

	go func() {
		l, err := net.Listen("tcp", getGRPCEndpoint())
		helper.PaincErr(err)
		s := grpc.NewServer(serverOpts...)
		defer s.GracefulStop()

		// 开启grpc客户端
		setup.RegisterGrpc(s)
		helper.PaincErr(s.Serve(l))
	}()

	// 开启 gateway http 客户端
	mux := runtime.NewServeMux()
	ctx := context.Background()

	grpcDialOptions := getGrpcDialOptions()

	conn, err := grpc.Dial(getGRPCEndpoint(), grpcDialOptions...)
	helper.PaincErr(err)
	setup.RegisterGateway(ctx, conn, mux)
	helper.PaincErr(Listen(mux))

}

func Listen(mux *runtime.ServeMux) error {
	if vars.HTTP_TLS_ON {
		return http.ListenAndServeTLS(getHttpGateWayEndpoint(), vars.CERTFILE, vars.KEYFILE, mux)
	}
	return http.ListenAndServe(getHttpGateWayEndpoint(), mux)
}

func getGrpcDialOptions() []grpc.DialOption {
	var grpcDialOptions []grpc.DialOption
	if vars.GRPC_TLS_ON {
		cred, err := credentials.NewClientTLSFromFile(vars.CERTFILE, "")
		helper.PaincErr(err)
		grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(cred))
	} else {
		grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return grpcDialOptions
}

func GetServerOpts() []grpc.ServerOption {
	var serverOpts []grpc.ServerOption
	if vars.GRPC_TLS_ON {
		serverOpts = utils.GetCred()
	} else {
		serverOpts = utils.GetInsecureCred()
	}
	return serverOpts
}

func getGRPCEndpoint() string {
	if vars.GRPC_TLS_ON {
		return ":" + vars.PORT_GRPC_TLS
	}
	return ":" + vars.PORT_GRPC
}

func getHttpGateWayEndpoint() string {
	if vars.HTTP_TLS_ON {
		return ":" + vars.PORT_GATEWAY_HTTP_TLS
	}
	return ":" + vars.PORT_GATEWAY_HTTP
}
