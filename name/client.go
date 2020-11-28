package main

import (
	"context"
	"log"

	gral "github.com/jamoreno22/lab2_dist/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	client := gral.NewNameNodeClient(conn)

	//Tiene que ser un stream pero no se cómo hacerlo :v
	//prop := gral.Proposal{Ip: "8000", Chunk: &gral.Chunk{Name: "Chunk1", Data: []byte("ABC€")}}
	proposals := []gral.Proposal{}
	proposals = append(proposals, gral.Proposal{Ip: "8000", Chunk: &gral.Chunk{Name: "Chunk1", Data: []byte("ABC€")}})
	proposals = append(proposals, gral.Proposal{Ip: "8001", Chunk: &gral.Chunk{Name: "Chunk2", Data: []byte("ABC2")}})

	writeLog2(client, proposals)
	//aber, err := client.WriteLog(context.Background(), prop)

	log.Println("fin")

}

//enviar propuesta al servidor
func writeLog2(nc gral.NameNodeClient, proposals []gral.Proposal) error {
	log.Println("Inicio de stream writeLog")
	stream, err := nc.WriteLog(context.Background())
	if err != nil {
		log.Printf("Error de stream writeLog: %v", err)
	}

	for _, prop := range proposals {
		if err := stream.Send(&prop); err != nil {
			log.Println("error al enviar prop")
			log.Fatalf("%v.Send() = %v", stream, err)
		}
		log.Printf("aqui voy-- prop.ip: %v", prop.Ip)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error recepcion response: %v", err)
	}
	log.Printf("Route summary: %v", reply)
	return nil
}
