package handler

import (
	"net/http"
	"strings"

	"github.com/DHBW-Malte/devops-lecture-project/checkout-service/internal/service"
	"github.com/DHBW-Malte/devops-lecture-project/pkg/httpx"
)

func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		httpx.Error(w, http.StatusUnauthorized, "Missing Authorization header")
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		httpx.Error(w, http.StatusUnauthorized, "Authorization header must use Bearer scheme")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
	if !service.VerifyToken(tokenString) {
		httpx.Error(w, http.StatusUnauthorized, "Invalid Token")
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"message": "Order placed successfully"})
}
