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
	// go SendSign(context.Background(), &pb.SendSignature{
	// 	Data: (message),
	// 	Signature: (signature),
	// })

	// return &pb.SignReply{Data: (message), Signature: (signature), Publickey: (public)}, nil

// }



	/*CHANGES */
// func (s *server) SendSign(ctx context.Context, in *pb.SignSignature) (*pb.SignReply, error) {
	const (
	address = "localhost:50052"
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewSignServiceClient(conn)
	/*params := pbc.GenerateA(160, 512)
  	pairing := params.NewPairing()
 	g := pairing.NewG2().Rand() */
	// message := "some text to sign"

	r, err := c.SendSign(context.Background(), &pb.SendSignature{

		SharedParams : pairing.String(),
		SharedG : g.Bytes(),
		Data : message,
  	Signature : signature,

  })
	// r1, err1 := c.SignString(context.Background(), &pb.SignRequest{
	//
	// 	SharedParams : in.SharedParams,
  // 		SharedG : in.SharedG,
	//
	// })

	return &pb.SignReply{Data: (message), Signature: (signature), Publickey: (public)}, nil

	/* END */



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
