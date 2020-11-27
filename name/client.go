package main

import (
	"context"
	"log"

	name "github.com/jamoreno22/lab2_dist/pkg/proto/NameNode"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	client := name.NewNameNodeClient(conn)
	//Tiene que ser un stream pero no se cómo hacerlo :v
	prop := name.Proposal{Ip: "8000", Chunk: &name.Chunk{Name: "Chunk1", Data: []byte("ABC€")}}

	aber, err := client.WriteLog(context.Background(), prop)

	log.Println(aber)

}
