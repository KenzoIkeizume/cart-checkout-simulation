package discount

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
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
	host := viper.GetString("discount_service_host")
	port := viper.GetString("discount_service_port")
	timeout := viper.GetInt("discount_service_timeout_seconds")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTimeout(time.Duration(timeout)*time.Second), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		fmt.Printf("did not connect: %v", err)
		// do not lock the process
		return nil, nil
	}

	c := discount_pc.NewDiscountClient(conn)

	return c, conn
}

func (cc discountService) GetDiscount(productID int32) float32 {
	discountClient, conn := cc.connection()
	if discountClient == nil || conn == nil {
		// do not lock the process
		return 0
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := discountClient.GetDiscount(ctx, &discount_pc.GetDiscountRequest{ProductID: productID})
	if err != nil {
		log.Printf("could not get discount: %v", err)
		// do not lock the process
		return 0
	}

	return r.GetPercentage()
}
