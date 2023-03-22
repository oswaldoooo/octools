package jwttoken

import (
	"math/rand"
	"time"
)

type JwtToken struct {
	secret      string
	Expiredtime time.Duration
	Issuer      string
	subject     string
	aud         string
	id          string
}

// detail set for jwt-token
func (s *JwtToken) DetailSet(subject, aud, id string) {
	s.aud = aud
	s.id = id
	s.subject = subject
}
func (s *JwtToken) SetScret(value string) {
	s.secret = value
}
func NewJwt() *JwtToken {
	jt := new(JwtToken)
	jt.Expiredtime = time.Duration(5) * time.Minute
	secret := ""
	for i := 0; i < 50; i++ {
		secret += string(byte(rand.Intn(26) + 97))
	}
	jt.SetScret(secret)
	return jt
}
