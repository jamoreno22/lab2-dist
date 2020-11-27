package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	data "github.com/jamoreno22/lab2_dist/pkg/proto/DataNode"
	name "github.com/jamoreno22/lab2_dist/pkg/proto/NameNode"
	"google.golang.org/grpc"
)

type dataServer struct {
	data.UnimplementedDataNodeServer
}

type nameServer struct {
	name.UnimplementedNameNodeServer
	Proposals map[string][]*name.Proposal
}

var path = "Log"

func main() {
	createFile()
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

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	sum := 0
	//for i , s:= range book.GetChunks(){
	//	fmt.Println(i,s)
	//	 sum = i
	//}
	_, err = file.WriteString(fmt.Sprint(sum))
}

func (s *nameServer) WriteLog(ctx context.Context, stream name.NameNode_WriteLogServer) (name.Message, error) {
	log.Printf("Stream WriteLogServer")
	for {
		prop, err := stream.Recv()
		if err == io.EOF {
			return name.Message{Text: "Oh no"}, err
		}
		if err != nil {
			return name.Message{Text: "Oh no 2"}, err
		}
		key := prop.Ip
		for _, n := range s.Proposals[key] {
			if err := stream.Send(*n); err != nil {
				return name.Message{Text: "Oh no 3"}, err
			}
		}

	}

	return name.Message{Text: "holi"}, nil

}
