package DataNode

import (
	"log"
	"net"

	"github.com/jamoreno22/lab2_dist/blob/main/DataNode/DataNode_grpc"

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
