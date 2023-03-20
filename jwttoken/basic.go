package jwttoken

import "math/rand"

type JwtToken struct {
	secret string
}

func (s *JwtToken) SetScret(value string) {
	s.secret = value
}
func NewJwt() *JwtToken {
	jt := new(JwtToken)
	secret := ""
	for i := 0; i < 12; i++ {
		secret += string(byte(rand.Intn(26) + 97))
	}
	jt.SetScret(secret)
	return jt
}
