package main

import (
	"context"
	"fmt"

	"github.com/canbefree/tools/helper"
	"github.com/org/repo/proto/pb_demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.TODO()
	conn, err := grpc.Dial("backend:8081", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUserAgent("http://proxy:3010"))
	helper.PaincErr(err)
	client := pb_demo.NewDemoServiceClient(conn)
	resp, err := client.ListDemos(ctx, &pb_demo.ListDemosRequest{})
	helper.PaincErr(err)
	fmt.Println(resp)
}

func testTls(){
	tls.NewCredentials
}
