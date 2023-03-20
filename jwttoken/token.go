package jwttoken

import (
	"oc_oauth/basic"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// defined the claims
type Custom_Claims struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// generate token

func GenerateToken(userid, username string) (token string, err error) {
	claims := Custom_Claims{
		UserId:   userid,
		UserName: username,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(5) * time.Minute))
	claims.Issuer = "brotherhood"
	token_origin := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = token_origin.SignedString([]byte(basic.Secret))
	return
}
