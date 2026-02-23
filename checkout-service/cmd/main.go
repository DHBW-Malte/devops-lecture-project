package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DHBW-Malte/devops-lecture-project/checkout-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler)

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
