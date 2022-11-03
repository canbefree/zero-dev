package utils

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Server interface {
	GRPCServer
	GateWayServer
}

type GRPCServer interface {
	StartServer() error
}

type GateWayServer interface {
	SetPort()
	Liten() error
}

type GRPC interface {
	Listen(mux *runtime.ServeMux) error
	SetGrpcDialOptions()
	SetGrpcEndPoints()
	SetGrpcServerOptions()
	SetHttpGatewayWayEndPoints()
}

// func NewGRPC() GRPC {
// }

// type Grpc struct {
// 	grpcEndPoints        string
// 	grpcOptions          []grpc.DialOption
// 	httpGatewayEndPoints string
// 	serverOptions        []grpc.ServerOption
// }

// type HttpGrpc struct {
// 	Grpc
// }

// func (h *HttpGrpc) Listen(mux *runtime.ServeMux) error {
// 	return http.ListenAndServe(getHttpGateWayEndpoint(), mux)
// }

// func (h *HttpGrpc) SetGrpcDialOptions() {
// 	h.Grpc.grpcOptions = append(h.Grpc.grpcOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
// }

// func (h *HttpGrpc) SetGrpcEndPoints() {
// 	h.grpcEndPoints = ":" + vars.PORT_GRPC
// }

// func (h *HttpGrpc) SetGrpcServerOptions() {
// 	h.serverOptions = utils.GetInsecureCred()
// }

// func (h *HttpGrpc) SetHttpGatewayWayEndPoints() {
// 	h.httpGatewayEndPoints = ":" + vars.PORT_GATEWAY_HTTP
// }

// type TlsGrpc struct {
// 	Grpc
// }
