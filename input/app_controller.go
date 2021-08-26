package input

import (
	"log"
	"net/http"

	cart_controller "cart-checkout-simulation/input/cart"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func AppController() {
	router := mux.NewRouter()
	cart_controller.SetRouter(router)
	router.HandleFunc("", notFound)

	println("Listen server on port 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}
