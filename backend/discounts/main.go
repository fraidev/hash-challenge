package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fraifelipe/hash-challenge/backend/discounts/protos/discount"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.DiscountServiceServer
}

func (s *server) GetDiscount(ctx context.Context, in *pb.DiscountRequest) (*pb.DiscountReply, error) {
	log.Printf("Received UserId: %v", in.GetUserId())

	

	// pg := pg.Connect(&pg.Options{
	// 	Addr:     app.config.DBAddr,
	// 	User:     app.config.DBUser,
	// 	Password: app.config.DBPassword,
	// 	Database: app.config.DBName,
	// })

	var products []*pb.ProductDiscount

	for _, product := range in.Products {
		log.Printf("Received ProductId: %v", product.Id)
		p := pb.ProductDiscount { ProductId: product.Id, ValueInCents: 1, Percentage: 5.0 };
		products = append(products, &p)
	}

	return &pb.DiscountReply{Products: products}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDiscountServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
