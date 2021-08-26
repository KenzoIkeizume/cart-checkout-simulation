package cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	application "cart-checkout-simulation/app"
	datastore "cart-checkout-simulation/infra/datastore"
	repository "cart-checkout-simulation/infra/repository"
	cart_request "cart-checkout-simulation/input/cart/request"
)

type cartController struct {
	cartApplication application.CartApplication
}

type CartController interface {
	post(w http.ResponseWriter, r *http.Request)
}

func NewCartController() CartController {
	db := datastore.NewDatabase()
	productRepository := repository.NewProductRespository(db)
	cartApplication := application.NewCartApplication(productRepository)
	return &cartController{cartApplication}
}

func (cc cartController) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var cr *cart_request.CartRequest
	cr = new(cart_request.CartRequest)

	err := decoder.Decode(&cr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "invalid payload"}`))
		return
	}

	cartResponse, err := cc.cartApplication.GetCart(cr)
	if err != nil {
		messageError := fmt.Sprintf(`{"message": "%s"}`, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(messageError))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartResponse)
}

func SetRouter(router *mux.Router) {
	cc := NewCartController()
	userRoute := router.PathPrefix("/carts").Subrouter()
	userRoute.HandleFunc("", cc.post).Methods(http.MethodPost)
}
