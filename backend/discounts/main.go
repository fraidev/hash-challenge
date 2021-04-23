package main

import (
	"context"
	"log"
	"net"

	models "github.com/fraifelipe/hash-challenge/backend/discounts/models"
	services "github.com/fraifelipe/hash-challenge/backend/discounts/services"
	pb "github.com/fraifelipe/hash-challenge/backend/discounts/protos/discount"
	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.DiscountServiceServer
}

func (s *server) GetDiscount(ctx context.Context, response *pb.DiscountRequest) (*pb.DiscountReply, error) {
	log.Printf("Received UserId: %v", response.GetUserId())

	//Connect database
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "YOUR_PASSWORD_HERE",
		Database: "hash",
	})

	// Map product ids from request
	var productsIds []uuid.UUID
	for _, p := range response.Products {
		id := uuid.FromStringOrNil(p.Id)
		productsIds = append(productsIds, id)
	}

	//Find product on database by ids
	products, err := FindProductByIds(db, productsIds)
	if err != nil {
		panic(err)
	}

	//Find user by user id header
	user, err := FindUserById(db, uuid.FromStringOrNil(response.GetUserId()))
	if user == nil {
		panic("invalid user")
	}
	if err != nil {
		panic(err)
	}

	log.Printf("User name: %v %v", user.FirstName, user.LastName)
	log.Printf("User id: %v", user.ID)
	log.Printf("User birth: %v", user.DateOfBirth)

	//Calculate discount
	var productsRespose []*pb.ProductDiscount
	for _, product := range products {
		log.Printf("Received ProductId: %v", product.ID)
		productDiscount := services.CalculateDiscount(product, user)
		productsRespose = append(productsRespose, &productDiscount)
	}

	return &pb.DiscountReply{Products: productsRespose}, nil
}

func FindProductByIds(db *pg.DB, ids []uuid.UUID) ([]*models.Product, error) {
	products := make([]*models.Product, len(ids))

	var err = db.Model(&products).WhereIn("id IN (?)", ids).Select()

	if err != nil {
		return nil, err
	}

	return products, nil
}


func FindUserById(db *pg.DB, id uuid.UUID) (*models.User, error) {
	users := make([]*models.User, 1)

	var err = db.Model(&users).Where("id = (?)", id).Select()

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
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
