package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	data "github.com/jamoreno22/lab2_dist/pkg/proto/DataNode"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	dc := data.NewDataNodeClient(conn)

	// -    - - - - - - -  - -    particionar pdf en chunks - - - - -  - - - -

	fileToBeChunked := "books/Mujercitas-Alcott_Louisa_May.pdf"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	book := make([]*data.Chunk, totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "somebigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		// books instantiation
		book[i] = &data.Chunk{Name: fileName, Data: partBuffer}

		fmt.Println("Split to : ", fileName)
		log.Println("tamaño: ", partSize)
	}
	// - - - - - --- -- - -  stream chunks - - - - - - - - - - - -

	// Create a random number of random points

	stream, err := dc.UploadBook(context.Background())
	if err != nil {
		log.Println("Error de stream uploadBook")
	}
	a := 1
	for _, chunk := range book {
		if err := stream.Send(chunk); err != nil {
			log.Println("error al enviar chunk")
			log.Fatalf("%v.Send(%d) = %v", stream, a, err)
		}
		a = a + 1
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error recepcion response")
	}
	log.Printf("Route summary: %v", reply)

	//dc.UploadBook(context.Background(), &book)

	log.Println("Client connected...")

}
