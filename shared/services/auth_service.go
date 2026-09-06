package services

import (
	"crypto/ed25519"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Ed25519Claims struct {
	ClientID string `json:"client_id"`
	jwt.RegisteredClaims
}

type AuthTokenService struct {
	db                *sql.DB
	privateKey        ed25519.PrivateKey
	publicKey         ed25519.PublicKey
	clientCredentials map[string]string
}

func NewAuthTokenService(db *sql.DB) (*AuthTokenService, error) {
	_ = godotenv.Load()

	privateKeyPEM := os.Getenv("PRIVATE_KEY_VALUE")
	if privateKeyPEM == "" {
		return nil, errors.New("PRIVATE_KEY_VALUE must be set in environment")
	}

	publicKeyPEM := os.Getenv("PUBLIC_KEY_VALUE")
	if publicKeyPEM == "" {
		return nil, errors.New("PUBLIC_KEY_VALUE must be set in environment")
	}

	privateKey, err := parseEd25519PrivateKey(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey, err := parseEd25519PublicKey(publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	clientID := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || secret == "" {
		return nil, errors.New("CLIENT_ID and CLIENT_SECRET must be set in environment")
	}

	clientCredentials := map[string]string{
		clientID: secret,
	}

	return &AuthTokenService{
		db:                db,
		privateKey:        privateKey,
		publicKey:         publicKey,
		clientCredentials: clientCredentials,
	}, nil
}

func parseEd25519PrivateKey(pemStr string) (ed25519.PrivateKey, error) {
	pemStr = strings.ReplaceAll(pemStr, "\\n", "\n")

	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to decode PEM block for private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	edKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("key is not an Ed25519 private key")
	}

	return edKey, nil
}

func parseEd25519PublicKey(pemStr string) (ed25519.PublicKey, error) {
	pemStr = strings.ReplaceAll(pemStr, "\\n", "\n")

	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to decode PEM block for public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKIX public key: %w", err)
	}

	edKey, ok := pub.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("key is not an Ed25519 public key")
	}

	return edKey, nil
}

func (s *AuthTokenService) GenerateToken(clientID string, clientSecret string) (string, string, string, error) {
	fmt.Printf("[INFO] Iniciando geração de token Ed25519 para client_id=%s\n", clientID)

	storedSecret, ok := s.clientCredentials[clientID]
	if !ok {
		fmt.Printf("[WARN] client_id não encontrado: %s\n", clientID)
		return "", "", "", errors.New("invalid client_id or secret")
	}
	if storedSecret != clientSecret {
		fmt.Printf("[WARN] Segredo inválido para client_id=%s\n", clientID)
		return "", "", "", errors.New("invalid client_id or secret")
	}

	expiration := time.Now().UTC().Add(2 * time.Hour)
	fmt.Printf("[DEBUG] Expiração definida para: %s\n", expiration.Format(time.RFC3339))

	claims := Ed25519Claims{
		ClientID: clientID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "ecommerce-api-go",
			Subject:   clientID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		fmt.Printf("[ERROR] Falha ao assinar token JWT com Ed25519: %v\n", err)
		return "", "", "", fmt.Errorf("error signing token: %w", err)
	}

	fmt.Printf("[INFO] Token JWT Ed25519 gerado com sucesso para client_id=%s\n", clientID)

	_, err = s.db.Exec(
		`INSERT INTO public.auth_tokens (client_id, jwt_token, expires_at) VALUES ($1, $2, $3)`,
		clientID,
		tokenString,
		expiration.UTC(),
	)
	if err != nil {
		fmt.Printf("[ERROR] Falha ao salvar token no banco de dados: %v\n", err)
		return "", "", "", fmt.Errorf("error saving to database: %w", err)
	}

	fmt.Printf("[INFO] Token salvo no banco com sucesso para client_id=%s\n", clientID)

	return tokenString, clientID, clientSecret, nil
}

func (s *AuthTokenService) ValidateToken(tokenString string) error {
	var exists bool
	err := s.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM public.auth_tokens WHERE jwt_token = $1 AND expires_at > NOW())`,
		tokenString,
	).Scan(&exists)
	if err != nil || !exists {
		return errors.New("invalid token or expired")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "EdDSA" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (s *AuthTokenService) GetValidToken(clientID string) (string, error) {
	var token string
	err := s.db.QueryRow(
		`SELECT jwt_token FROM public.auth_tokens WHERE client_id = $1 AND expires_at > NOW() ORDER BY created_at DESC LIMIT 1`,
		clientID,
	).Scan(&token)

	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return token, nil
}
