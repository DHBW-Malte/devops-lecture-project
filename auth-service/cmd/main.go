package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DHBW-Malte/devops-lecture-project/auth-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", handler.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", handler.AuthLogoutHandler)

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d"), mux))
}
