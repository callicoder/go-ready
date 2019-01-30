package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenService struct {
	config config.AuthConfig
}

func NewTokenService(config config.AuthConfig) *TokenService {
	return &TokenService{
		config: config,
	}
}

func (s *TokenService) CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Subject:   string(user.Id),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(s.config.JwtExpiryInSec) * time.Second).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(s.config.JwtSecret))

	if err != nil {
		return "", err
	}

	return tokenStr, err
}

func (s *TokenService) ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.config.JwtSecret), nil
	})

	return token, err
}

func (s *TokenService) GetUserSessionFromToken(tokenStr string) (*model.Session, error) {
	token, err := s.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		userId, _ := strconv.ParseInt(claims.Subject, 10, 64)

		session := &model.Session{
			UserId: int(userId),
		}

		return session, nil
	}

	return nil, err
}
