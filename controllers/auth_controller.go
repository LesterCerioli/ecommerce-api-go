package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"ecommerce-api-go/shared/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ValidateRequest struct {
	Token string `json:"token"`
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// GenerateToken godoc
// @Summary      Gera um novo token JWT
// @Description  Cria um token JWT usando client_id e client_secret
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body TokenRequest true "Credenciais do cliente"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Router       /auth/token [post]
func (c *AuthController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.AuthService.GenerateToken(context.Background(), req.ClientID, req.ClientSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// ValidateToken godoc
// @Summary      Valida um token JWT
// @Description  Verifica se o token JWT é válido e não expirou
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body ValidateRequest true "Token JWT"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Router       /auth/token/validate [post]
func (c *AuthController) ValidateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.AuthService.ValidateToken(context.Background(), req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
