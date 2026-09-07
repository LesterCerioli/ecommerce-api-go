package services

import (
	"context"
	"crypto/hmac"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const TokenTTLSeconds = 120

type AuthService struct {
	db           *sql.DB
	clientID     string
	clientSecret string
	privateKey   []byte
	publicKey    []byte
}

func NewAuthService(dsn string) (*AuthService, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	privateKey := []byte(os.Getenv("PRIVATE_KEY_VALUE"))
	publicKey := []byte(os.Getenv("PUBLIC_KEY_VALUE"))

	return &AuthService{
		db:           db,
		clientID:     os.Getenv("CLIENT_ID"),
		clientSecret: os.Getenv("CLIENT_SECRET"),
		privateKey:   privateKey,
		publicKey:    publicKey,
	}, nil
}

func (a *AuthService) EnsureTable(ctx context.Context) error {
	ddl := `
    CREATE TABLE IF NOT EXISTS public.auth_tokens (
        id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        client_id   VARCHAR(64) NOT NULL,
        jwt_token   TEXT NOT NULL,
        created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        expires_at  TIMESTAMPTZ NOT NULL
    );`
	_, err := a.db.ExecContext(ctx, ddl)
	return err
}

func (a *AuthService) authenticateCredentials(clientID, clientSecret string) bool {
	return hmac.Equal([]byte(clientID), []byte(a.clientID)) &&
		hmac.Equal([]byte(clientSecret), []byte(a.clientSecret))
}

func (a *AuthService) createJWT(clientID string, tokenID uuid.UUID) (string, error) {
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":  clientID,
		"iat":  now.Unix(),
		"exp":  now.Add(time.Second * TokenTTLSeconds).Unix(),
		"iss":  "auth_service",
		"type": "m2m",
		"jti":  tokenID.String(),
	}

	// Ed25519 (EdDSA) assinatura
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(a.privateKey)
}

func (a *AuthService) verifyJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (a *AuthService) GenerateToken(ctx context.Context, clientID, clientSecret string) (map[string]interface{}, error) {
	if !a.authenticateCredentials(clientID, clientSecret) {
		return nil, errors.New("invalid client_id or client_secret")
	}

	now := time.Now().UTC()
	expiresAt := now.Add(time.Second * TokenTTLSeconds)
	tokenID := uuid.New()

	jwtToken, err := a.createJWT(clientID, tokenID)
	if err != nil {
		return nil, err
	}

	_, err = a.db.ExecContext(ctx,
		`INSERT INTO public.auth_tokens (id, client_id, jwt_token, created_at, expires_at)
         VALUES ($1, $2, $3, $4, $5)`,
		tokenID, clientID, jwtToken, now, expiresAt)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"access_token": jwtToken,
		"token_type":   "Bearer",
		"expires_in":   TokenTTLSeconds,
		"expires_at":   expiresAt.Format(time.RFC3339),
	}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, jwtToken string) (map[string]interface{}, error) {
	claims, err := a.verifyJWT(jwtToken)
	if err != nil {
		return nil, err
	}

	clientID := claims["sub"].(string)

	row := a.db.QueryRowContext(ctx,
		`SELECT id, client_id, expires_at
         FROM public.auth_tokens
         WHERE jwt_token = $1 AND client_id = $2
         ORDER BY created_at DESC
         LIMIT 1`,
		jwtToken, clientID)

	var id uuid.UUID
	var dbClientID string
	var expiresAt time.Time
	if err := row.Scan(&id, &dbClientID, &expiresAt); err != nil {
		return nil, errors.New("token not found in database")
	}

	if expiresAt.Before(time.Now().UTC()) {
		return nil, errors.New("token has expired in database")
	}

	return map[string]interface{}{
		"valid":      true,
		"client_id":  dbClientID,
		"token_id":   id.String(),
		"expires_at": expiresAt.Format(time.RFC3339),
	}, nil
}

func (a *AuthService) RevokeToken(ctx context.Context, jwtToken string) (bool, error) {
	res, err := a.db.ExecContext(ctx,
		`DELETE FROM public.auth_tokens WHERE jwt_token = $1`, jwtToken)
	if err != nil {
		return false, err
	}
	rows, _ := res.RowsAffected()
	return rows > 0, nil
}
