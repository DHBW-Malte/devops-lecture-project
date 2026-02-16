package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func main() {
	mux := http.NewServeMux()
	// Auth Service
	mux.HandleFunc("/auth/login", authLoginHandler)
	mux.HandleFunc("/auth/logout", authLogoutHandler)
	// Product Service
	mux.HandleFunc("/products", productListHandler)
	mux.HandleFunc("/products/{id}", productDetailHandler)
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	username := r.FormValue("username")
	password := r.FormValue("password")
	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := createToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error generating the token"}`))
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid credentials"}`))
	}
}

func authLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// In this simple example, we'll just return a success message
	w.Write([]byte(`{"message": "Logout successful"}`))
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

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	return err == nil && token.Valid
}
