package main

import (
	"fmt"
	"log"
	"net"

	data "github.com/jamoreno22/lab2_dist/pkg/proto/DataNode"
	name "github.com/jamoreno22/lab2_dist/pkg/proto/NameNode"
	"google.golang.org/grpc"
)

type dataServer struct {
	data.UnimplementedDataNodeServer
}

type nameServer struct {
	name.UnimplementedNameNodeServer
}

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} // create a server instance
	ds := dataServer{}                               // create a gRPC server object
	grpcDataServer := grpc.NewServer()               // attach the Ping service to the server
	data.RegisterDataNodeServer(grpcDataServer, &ds) // start the server

	log.Println("Server running ...")
	if err := grpcDataServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	// create a listener on TCP port 7777
	namelis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} // create a server instance
	ns := nameServer{}                               // create a gRPC server object
	grpcNameServer := grpc.NewServer()               // attach the Ping service to the server
	name.RegisterNameNodeServer(grpcNameServer, &ns) // start the server

	log.Println("Server running ...")
	if err := grpcNameServer.Serve(namelis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
