package input

import (
	"fmt"
	"log"
	"net/http"

	cart_controller "cart-checkout-simulation/input/cart"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	port := viper.GetString("app_port")
	println("Listen server on port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
