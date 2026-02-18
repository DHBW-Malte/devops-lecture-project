package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/DHBW-Malte/devops-lecture-project/internal/handler"
)

var secretKey = []byte("secret-key")

func main() {
	mux := http.NewServeMux()
	// Auth Service
	mux.HandleFunc("/auth/login", handler.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", handler.AuthLogoutHandler)
	// Product Service
	mux.HandleFunc("/products", handler.ProductListHandler)
	mux.HandleFunc("/products/{id}", handler.ProductListHandler)
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func checkoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Missing Authorization header"}`))
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Authorization header must use Bearer scheme"}`))
		return
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

	if !verifyToken(tokenString) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Invalid token"}`))
		return
	}

	w.Write([]byte(`{"message":"Order placed successfully"}`))
}
