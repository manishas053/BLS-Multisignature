package main

import (
	"log"
	"net"

	pb "google.golang.org/grpc/examples/bls/bls"
	"github.com/Nik-U/pbc"
	"crypto/sha256"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SignString(ctx context.Context, in *pb.SignRequest) (*pb.SignReply, error) {

	pairing, _ := pbc.NewPairingFromString(in.SharedParams)
	g := pairing.NewG2().SetBytes(in.SharedG)

	privateKey := pairing.NewZr().Rand()
  	publicKey := pairing.NewG2().PowZn(g, privateKey)
	public := publicKey.Bytes()

	message := "some text to sign"
  	h := pairing.NewG1().SetFromStringHash(message, sha256.New())

	sign := pairing.NewG2().PowZn(h, privateKey)
	signature := sign.Bytes()

	return &pb.SignReply{Data: (message), Signature: (signature), Publickey: (public)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSignServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
