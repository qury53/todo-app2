package auth_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/qury53/todo-app2/internal/core/domain"
)

type AuthService struct {
	AuthRepository  AuthRepository
	JWTSecretKey    []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error)
}

func NewAuthService(config *Config, repository AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository:  repository,
		JWTSecretKey:    []byte(config.JWTSecretKey),
		AccessTokenTTL:  config.AccessTokenTTL,
		RefreshTokenTTL: config.RefrshTokenTTL,
	}
}

func (s *AuthService) GetRefreshTokenTTL() time.Duration {
	return s.RefreshTokenTTL
}

func (s *AuthService) GenerateAccessToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "todoapp",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) GenerateRefreshToken(userID int) (string, error) {
	expirationTime := time.Now().Add(s.RefreshTokenTTL)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprintf("%d", userID),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.JWTSecretKey)
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return s.JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (s *AuthService) ValidateRefreshToken(ctx context.Context, tokenStr string) (int, bool, error) {
	claims := &Claims{}

	// 1. Парсим токен с заполнением нашей структуры claims
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Обязательно проверяем метод подписи (чтобы избежать атаки с подменой алгоритма на "none")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.JWTSecretKey, nil
	})

	// 2. Обрабатываем ошибки парсинга (например, токен просрочен или поврежден)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			// Токен просрочен — это не системная ошибка, а штатный сценарий авторизации
			return 0, false, nil
		}
		// Любая другая ошибка (неверная подпись, некорректный JSON и т.д.)
		return 0, false, fmt.Errorf("failed to parse jwt: %w", err)
	}

	// 3. Дополнительная проверка на валидность токена
	if !token.Valid {
		return 0, false, nil
	}

	// 4. Всё отлично — возвращаем ID пользователя и подтверждаем валидность
	return claims.UserID, true, nil
}
