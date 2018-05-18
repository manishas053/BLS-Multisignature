package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Nik-U/pbc"
	"crypto/sha256"

	pb "github.com/bithinalangot/bls/list"
)

const (
	address     = "localhost:50051"
	defaultName = "list"
)

func insertNode(client pb.ListClient, node *pb.NodeRequest) {
	resp, err := client.InsertNode(context.Background(), node)
	if err != nil {
		log.Fatalf("node couldn't insert %v", err)
	}
	if resp.Success {
		log.Printf("New Node inserted %v", node.Data)
	}
}

func printing(client pb.ListClient, node *pb.LinkRequest) {
	stream, err := client.Printing(context.Background(), node)
	if err != nil {
		log.Fatalf("node couldn't not be retrieved %v", err)
	}
	for {
		node1, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't get the node")
		}
		log.Printf("Customer: %v", node1)
	}
}

func adding(client pb.ListClient, in *pb.InputMsg) {
	resp, err := client.ProcessInput(context.Background(), in)
	if err != nil {
		log.Fatalf("node couldn't insert %v", err)
	}
	if resp.Resp != 0 {
		log.Printf("New Node inserted %v", in.Data)
	}
}

func main() {
	// Setup connection to gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Didn't connect %v", err)
	}
	defer conn.Close()
	client := pb.NewListClient(conn)

	//node := &pb.NodeRequest{
	//	Data: 1,
	//}
	//insertNode(client, node)
	//node = &pb.NodeRequest{
	//	Data: 2,
	//}
	//insertNode(client, node)
	//nodes := &pb.LinkRequest{}
	//printing(client, nodes)

	params := pbc.GenerateA(160, 512)
  pairing := params.NewPairing()
 	g := pairing.NewG2().Rand()

	in := &pb.InputMsg{SharedParams : params.String(), SharedG : g.Bytes(), }
	adding(client, in)
}
