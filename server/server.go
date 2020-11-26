package main

import (
	"fmt"
	"log"
	"net"

	lab2 "github.com/jamoreno22/lab2_dist/pkg/proto/DataNode"
	"google.golang.org/grpc"
)

type server struct {
	lab2.UnimplementedDataNodeServer
}

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} // create a server instance
	s := server{}                               // create a gRPC server object
	grpcServer := grpc.NewServer()              // attach the Ping service to the server
	lab2.RegisterDataNodeServer(grpcServer, &s) // start the server

	log.Println("Server running ...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
