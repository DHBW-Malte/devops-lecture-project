package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DHBW-Malte/devops-lecture-project/products-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	// Product Service
	mux.HandleFunc("POST /products", handler.AddProduct)
	mux.HandleFunc("/products", handler.ProductListHandler)
	mux.HandleFunc("/products/{id}", handler.ProductDetailHandler)
	mux.HandleFunc("POST /products/filter", handler.FilterProductHandler)

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
