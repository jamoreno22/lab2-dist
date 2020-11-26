package main

import (
	"log"
	"net"

	"DataNode_grpc"

	"google.golang.org/grpc"
)

type server struct {
	DataNode_grpc.UnimplementedDataNodeServer
}

func main() {
	log.Printf("DataNode Server Running...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatallf(err)
	}

	srv := grpc.NewServer()
	DataNode_grpc.RegisterDataNodeServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}
