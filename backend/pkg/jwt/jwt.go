package jwt

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secret []byte
	expiration time.Duration
	log *slog.Logger
}

func NewJWTService(secret string, hours int, log *slog.Logger) *JWTService {
	return &JWTService{
		secret: []byte(secret),
		expiration: time.Duration(hours) * time.Hour,
		log: log,
	}
}

func (s *JWTService) CreateJWT(userID int64, username string) (string, error) {
	s.log.Info(fmt.Sprintf("secret info: %s", s.secret))
	claims:= &JWTClaim{
		UserID: userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Subject: fmt.Sprintf("%d", userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JWTService) ParseJWT(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok  := token.Claims.(*JWTClaim)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}
	return claims, nil
}