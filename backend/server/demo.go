package server

import (
	"context"

	"github.com/org/repo/proto/pb_demo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DemoServer struct {
	pb_demo.UnimplementedDemoServiceServer
}

func NewDemoServer() *DemoServer {
	return &DemoServer{}
}

func (s *DemoServer) ListDemos(_ context.Context, _ *pb_demo.ListDemosRequest) (*pb_demo.ListDemosResponse, error) {
	return &pb_demo.ListDemosResponse{
		Demos:         []*pb_demo.Demo{},
		NextPageToken: "hahah",
	}, nil
}

func (s *DemoServer) GetDemo(_ context.Context, _ *pb_demo.GetDemoRequest) (*pb_demo.Demo, error) {
	panic("not implemented") // TODO: Implement
}

func (s *DemoServer) CreateDemo(_ context.Context, _ *pb_demo.CreateDemoRequest) (*pb_demo.Demo, error) {
	panic("not implemented") // TODO: Implement
}

func (s *DemoServer) UpdateDemo(_ context.Context, _ *pb_demo.UpdateDemoRequest) (*pb_demo.Demo, error) {
	panic("not implemented") // TODO: Implement
}

func (s *DemoServer) DeleteDemo(_ context.Context, _ *pb_demo.DeleteDemoRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (s *DemoServer) mustEmbedUnimplementedDemoServiceServer() {
	panic("not implemented") // TODO: Implement
}
