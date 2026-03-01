package handler

import (
	"github.com/DHBW-Malte/devops-lecture-project/auth-service/internal/service"
	"github.com/DHBW-Malte/devops-lecture-project/pkg/httpx"
	"net/http"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "user" || password != "pass" {
		httpx.Error(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := service.CreateToken(username)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error generating the token")
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"token": token})
}

func AuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	httpx.JSON(w, http.StatusOK, map[string]string{"message": "Logout successful"})
}
