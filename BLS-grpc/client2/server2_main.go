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
	port = ":50052"
)

type server struct{}

func (s *server) SendSign(ctx context.Context, in2 *pb.SendSignature) (*pb.SignReply, error) {

	pairing, _ := pbc.NewPairingFromString(in2.SharedParams)
	g := pairing.NewG2().SetBytes(in2.SharedG)

	message := in2.Data
	//sign_recvd := in.Signature
	sign_recvd := pairing.NewG1().SetBytes(in2.Signature)

	privateKey := pairing.NewZr().Rand()
  publicKey := pairing.NewG2().PowZn(g, privateKey)
	public := publicKey.Bytes()

	//message := "some text to sign"
  h := pairing.NewG1().SetFromStringHash(message, sha256.New())

	sign := pairing.NewG2().PowZn(h, privateKey)
	//signature := sign.Bytes()

	aggregate := pairing.NewG2().Add(sign_recvd, sign)
	aggregate_sign := aggregate.Bytes()
	return &pb.SignReply{Data: (message), Signature: (aggregate_sign), Publickey: (public)}, nil
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
