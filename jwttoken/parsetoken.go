package jwttoken

import (
	"errors"
	"oc_oauth/basic"

	"github.com/golang-jwt/jwt/v5"
)

// parse the token
func ParseToken(tokenString string) (claims *Custom_Claims, err error) {
	var mc = new(Custom_Claims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) { return []byte(basic.Secret), nil })
	if err == nil {
		if token.Valid {
			return mc, nil
		} else {
			err = errors.New("invalid token")
		}
	}
	return
}
