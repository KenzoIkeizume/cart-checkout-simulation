package app

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	datastore "cart-checkout-simulation/infra/datastore"
	repository "cart-checkout-simulation/infra/repository"
	cart_request "cart-checkout-simulation/input/cart/request"
	cart_response "cart-checkout-simulation/input/cart/response"
	product_request "cart-checkout-simulation/input/product/request"
	product_response "cart-checkout-simulation/input/product/response"
)

type MyMockedDiscountObject struct {
	mock.Mock
}

func (m *MyMockedDiscountObject) GetDiscount(productID int32) float32 {
	if productID%2 == 0 {
		return 0
	}

	return 0.5
}

func TestMain(t *testing.T) {
	viper.Set("cents_math", 10000)
	viper.Set("blackfriday_day", 0)   // never will be blackfriday
	viper.Set("blackfriday_month", 0) // never will be blackfriday
}

func TestGetCart_Success(t *testing.T) {
	// database is in memory
	db := datastore.NewDatabase()
	// repo is already mock
	repo := repository.NewProductRespository(db)
	mockDiscountService := new(MyMockedDiscountObject)

	mockDiscountService.On("GetDiscount", 1).Return(0.5)

	app := NewCartApplication(repo, mockDiscountService)

	request := &cart_request.CartRequest{
		Products: []product_request.ProductRequest{
			{
				ID:       1,
				Quantity: 2,
			},
		},
	}

	response, err := app.GetCart(request)

	expectedResponse := cart_response.CartResponse{
		TotalAmount:             303140000,
		TotalAmountWithDiscount: 303135000,
		TotalDiscount:           5000,
		Products: []product_response.ProductResponse{
			{
				ID:          1,
				Quantity:    2,
				UnityAmount: 151570000,
				TotalAmount: 303140000,
				Discount:    5000,
				IsGift:      false,
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestGetCart_SuccessWithoutDiscount(t *testing.T) {
	// database is in memory
	db := datastore.NewDatabase()
	// repo is already mock
	repo := repository.NewProductRespository(db)
	mockDiscountService := new(MyMockedDiscountObject)

	mockDiscountService.On("GetDiscount", 1).Return(0.5)

	app := NewCartApplication(repo, mockDiscountService)

	request := &cart_request.CartRequest{
		Products: []product_request.ProductRequest{
			{
				ID:       2,
				Quantity: 2,
			},
		},
	}

	response, err := app.GetCart(request)

	expectedResponse := cart_response.CartResponse{
		TotalAmount:             1876220000,
		TotalAmountWithDiscount: 1876220000,
		TotalDiscount:           0,
		Products: []product_response.ProductResponse{
			{
				ID:          2,
				Quantity:    2,
				UnityAmount: 938110000,
				TotalAmount: 1876220000,
				Discount:    0,
				IsGift:      false,
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestGetCart_SuccessWithGift(t *testing.T) {
	_, m, d := time.Now().Date()

	viper.Set("blackfriday_day", d)
	viper.Set("blackfriday_month", int(m))

	// database is in memory
	db := datastore.NewDatabase()
	// repo is already mock
	repo := repository.NewProductRespository(db)
	mockDiscountService := new(MyMockedDiscountObject)

	mockDiscountService.On("GetDiscount", 1).Return(0.5)

	app := NewCartApplication(repo, mockDiscountService)

	request := &cart_request.CartRequest{
		Products: []product_request.ProductRequest{
			{
				ID:       2,
				Quantity: 2,
			},
		},
	}

	response, err := app.GetCart(request)

	expectedResponse := cart_response.CartResponse{
		TotalAmount:             1876220000,
		TotalAmountWithDiscount: 1876220000,
		TotalDiscount:           0,
		Products: []product_response.ProductResponse{
			{
				ID:          2,
				Quantity:    2,
				UnityAmount: 938110000,
				TotalAmount: 1876220000,
				Discount:    0,
				IsGift:      false,
			},
			{
				ID:          6,
				Quantity:    1,
				UnityAmount: 0,
				TotalAmount: 0,
				Discount:    0,
				IsGift:      true,
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, response)
}
