package utils

import (
	"github.com/canbefree/tools/helper"
	"github.com/org/bzero/vars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var GetTransportCredentials = func() credentials.TransportCredentials {
	c, err := credentials.NewServerTLSFromFile(vars.CERTFILE, vars.KEYFILE)
	helper.PaincErr(err)
	return c
}

var GetCred = func() []grpc.ServerOption {
	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.Creds(GetTransportCredentials()))
	return serverOpts
}

var GetInsecureCred = func() []grpc.ServerOption {
	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.EmptyServerOption{})
	return serverOpts
}
