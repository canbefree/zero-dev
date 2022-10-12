package main

import (
	"github.com/canbefree/tools/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var GetTransportCredentials = func() credentials.TransportCredentials {
	c, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	helper.PaincErr(err)
	return c
}

var GetCred = func() []grpc.ServerOption {
	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.Creds(GetTransportCredentials()))
	return serverOpts
}
