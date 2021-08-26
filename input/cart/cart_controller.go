package cart

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type cartController struct {
}

type CartController interface {
	post(w http.ResponseWriter, r *http.Request)
}

func NewCartController() CartController {
	return &cartController{}
}

func (cc cartController) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var cr *CartRequest
	cr = new(CartRequest)

	err := decoder.Decode(&cr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "invalid payload"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cr)
}

func SetRouter(router *mux.Router) {
	cc := NewCartController()
	userRoute := router.PathPrefix("/carts").Subrouter()
	userRoute.HandleFunc("", cc.post).Methods(http.MethodPost)
}
