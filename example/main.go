package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/king526/grpc2h"
	"github.com/king526/grpc2h/example/pb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Get(ctx context.Context, req *pb.GetItemReq) (*pb.Item, error) {
	return &pb.Item{
		Id:   req.Id,
		Name: "example_name",
	}, nil
}

func (s *server) List(_ *pb.Null, stream pb.ItemService_ListServer) error {
	stream.Send(&pb.Item{
		Id:   1,
		Name: "example1",
	})
	return stream.Send(&pb.Item{
		Id:   2,
		Name: "example2",
	})
}

func main() {
	// start grpc server
	var srv server
	gs := grpc.NewServer()
	ls, _ := net.Listen("tcp", "localhost:8080")
	pb.RegisterItemServiceServer(gs, &srv)
	go gs.Serve(ls)

	// now make a http server
	var b grpc2h.HandlerBuilder
	b.SetGRPCServer(srv)
	b.RegisterSeverStream(itemServiceListServer{})
	h, e := b.Build()
	if e != nil {
		fmt.Println(e)
		return
	}
	// h:=pb.GetHTTPHandler(&srv)
	http.ListenAndServe("localhost:8081", h)
}

// copy from pb.go
type itemServiceListServer struct {
	grpc.ServerStream
}

func (x *itemServiceListServer) Send(m *pb.Item) error {
	return x.ServerStream.SendMsg(m)
}
