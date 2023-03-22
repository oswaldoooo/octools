package jwttoken

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// defined the claims
type Custom_Claims struct {
	Args map[string]string `json:"args"`
	jwt.RegisteredClaims
}

// generate token

func (s *JwtToken) GenerateToken(args map[string]string) (token string, err error) {
	claims := Custom_Claims{
		Args: make(map[string]string),
	}
	claims.Args = args
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.Expiredtime))
	claims.Issuer = s.Issuer
	token_origin := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = token_origin.SignedString([]byte(s.secret))
	return
}
