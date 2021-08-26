package discount

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	discount_pc "cart-checkout-simulation/infra/discount/proto"
)

type discountService struct{}

type DiscountService interface {
	GetDiscount(productID int32) float32
}

func NewDiscountService() DiscountService {
	return &discountService{}
}

func (cc discountService) connection() (discount_pc.DiscountClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := discount_pc.NewDiscountClient(conn)

	return c, conn
}

func (cc discountService) GetDiscount(productID int32) float32 {
	discountClient, conn := cc.connection()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := discountClient.GetDiscount(ctx, &discount_pc.GetDiscountRequest{ProductID: productID})
	if err != nil {
		log.Fatalf("could not get discount: %v", err)
	}

	return r.GetPercentage()
}
