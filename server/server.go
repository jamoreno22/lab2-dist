package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	gral "github.com/jamoreno22/lab2_dist/pkg/proto"
	"google.golang.org/grpc"
)

type dataServer struct {
	gral.UnimplementedDataNodeServer
}

type nameServer struct {
	gral.UnimplementedNameNodeServer
	Proposals map[string][]*gral.Proposal
}

type nameServer2 struct {
	gral.UnimplementedNameNodeServer
}

var path = "Log"

// books variable when books are saved
var books = []gral.Book{}

func main() {
	/*
		// create a listener on TCP port 7777
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		// create a server instance
		ds := dataServer{}                               // create a gRPC server object
		grpcDataServer := grpc.NewServer()               // attach the Ping service to the server
		gral.RegisterDataNodeServer(grpcDataServer, &ds) // start the server

		log.Println("Server running ...")
		if err := grpcDataServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	*/

	// create a listener on TCP port 8000
	namelis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	ns := nameServer2{}                              // create a gRPC server object
	grpcNameServer := grpc.NewServer()               // attach the Ping service to the server
	gral.RegisterNameNodeServer(grpcNameServer, &ns) // start the server

	log.Println("NameServer running ...")
	if err := grpcNameServer.Serve(namelis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// UploadBook server side
func (d *dataServer) UploadBook(ubs gral.DataNode_UploadBookServer) error {
	log.Printf("Stream UploadBook")

	// saved Proposals array
	book := gral.Book{}
	indice := 0
	for {
		chunk, err := ubs.Recv()
		if err == io.EOF {
			books = append(books, book)
			log.Printf("EOF... books lenght = %d", len(books))
			return (ubs.SendAndClose(&gral.Message{Text: "EOF"}))
		}
		if err != nil {
			return err
		}
		book.Chunks = append(book.Chunks, chunk)
		indice = indice + 1

	}
}

// Writelog server side
func (s *nameServer) WriteLog(wls gral.NameNode_WriteLogServer) error {
	log.Printf("Stream WriteLogServer")
	// create log
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// saved Proposals array
	sP := []gral.Proposal{}

	for {
		prop, err := wls.Recv()
		if err == io.EOF {
			log.Printf("EOF ------------")
			return (wls.SendAndClose(&gral.Message{Text: "Oh no... EOF"}))
		}
		if err != nil {
			return err
		}

		sP = append(sP, *prop)
		log.Println("algo hace")
		// Aquí va el código para guardar el log

		_, err2 := f.WriteString("ip " + prop.Ip)

		if err2 != nil {
			log.Fatal(err2)
		}

		// ------------------------------------
	}
}
