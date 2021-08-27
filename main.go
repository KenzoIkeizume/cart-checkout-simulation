package main

import (
	"log"

	"github.com/spf13/viper"

	app_controller "cart-checkout-simulation/input"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	app_controller.AppController()
}
