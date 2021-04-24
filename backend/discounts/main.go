package main

import (
	"context"
	"log"
	"net"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"github.com/fraifelipe/hash-challenge/backend/discounts/infrastructure"
	"github.com/fraifelipe/hash-challenge/backend/discounts/models"
	"github.com/fraifelipe/hash-challenge/backend/discounts/repositories"
	"github.com/fraifelipe/hash-challenge/backend/discounts/services"

	grpc_discount "github.com/fraifelipe/hash-challenge/backend/discounts/protos/discount"
)

type server struct {
	grpc_discount.DiscountServiceServer
}

func (s *server) GetDiscount(ctx context.Context, response *grpc_discount.DiscountRequest) (*grpc_discount.DiscountReply, error) {
	userId := response.GetUserId()
	log.Printf("Received UserId: %v", userId)

	//Connect database
	db := infrastructure.GetDatabasConnection()

	// Map product ids from request
	productsIds := MapProductIdsFromRequest(response.Products)

	//Find product on database by ids
	products, err := repositories.NewProductRepository(db).FindProductByIds(productsIds)
	if err != nil {
		log.Printf(err.Error())
		return &grpc_discount.DiscountReply{}, nil
	}

	//Find user on database by user id header
	user, err := repositories.NewUserRepository(db).FindUserById(db, uuid.FromStringOrNil(userId))
	if user == nil {
		return &grpc_discount.DiscountReply{}, nil
	}
	if err != nil {
		log.Printf(err.Error())
		return &grpc_discount.DiscountReply{}, nil
	}
	log.Printf("User name: %v %v User id: %v User birth: %v", user.FirstName, user.LastName, user.ID, user.DateOfBirth)

	//Calculate product discounts
	productsRespose := CalculateAllProductDiscounts(products, user)

	return &grpc_discount.DiscountReply{Products: productsRespose}, nil
}

func CalculateAllProductDiscounts(products []*models.Product, user *models.User) []*grpc_discount.ProductDiscount {
	var productsRespose []*grpc_discount.ProductDiscount
	for _, product := range products {
		log.Printf("Received ProductId: %v", product.ID)
		productDiscount := services.CalculateDiscount(product, user)

		log.Printf("Percentage %v: Final Value: %v", productDiscount.Percentage, productDiscount.ValueInCents)
		productsRespose = append(productsRespose, &productDiscount)
	}

	return productsRespose
}

func MapProductIdsFromRequest(products []*grpc_discount.Product) []uuid.UUID {
	var productsIds []uuid.UUID
	for _, p := range products {
		id := uuid.FromStringOrNil(p.Id)
		productsIds = append(productsIds, id)
	}

	return productsIds
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_discount.RegisterDiscountServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
