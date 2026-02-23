package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DHBW-Malte/devops-lecture-project/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	// Auth Service
	mux.HandleFunc("/auth/login", handler.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", handler.AuthLogoutHandler)
	// Product Service
	mux.HandleFunc("/products", handler.ProductListHandler)
	mux.HandleFunc("/products/{id}", handler.ProductDetailHandler)
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler)

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
