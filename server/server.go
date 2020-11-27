package main

import (

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

type nameServer2 struct {
	name.UnimplementedNameNodeServer
}

var path = "Log"

func main() {
	//createFile()
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

	// create a listener on TCP port 8000
	namelis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} // create a server instance
	ns := nameServer2{}                               // create a gRPC server object
	grpcNameServer := grpc.NewServer()               // attach the Ping service to the server
	name.RegisterNameNodeServer(grpcNameServer, &ns) // start the server

	log.Println("Server running ...")
	if err := grpcNameServer.Serve(namelis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

/*
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


func writeFile(name.Proposal) {
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
*/
func (s *nameServer) WriteLog(wls name.NameNode_WriteLogServer) error {
	log.Printf("Stream WriteLogServer")
	// create log
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// saved Proposals array
	sP := []name.Proposal{}
	for {
		prop, err := wls.Recv()
		if err == io.EOF {
			log.Printf("EOF ------------")
			return (wls.SendAndClose(&name.Message{Text:"Oh no... EOF",}))
		}
		if err != nil {
			return err
		}

		sP = append(sP, *prop)
		
		// Aquí va el código para guardar el log		
		
		_, err2 := f.WriteString("ip " + prop.Ip)
		
		if err2 != nil {
			log.Fatal(err2)
		}

		// ------------------------------------
	}

}
