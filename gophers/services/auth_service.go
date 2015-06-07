package services

import (
	"crypto/sha512"
	"fmt"
)

type AuthService interface {
	Auth(username, password string) bool
}

type DummyAuthService struct {
}

func (as DummyAuthService) Auth(username, password string) bool {
	return true
}

type BasicAuthService struct {
	UpHashed map[string]string
}

func (as BasicAuthService) Auth(username, password string) bool {
	p, exists := as.UpHashed[username]
	if !exists {
		return false
	}
	sha := sha512.Sum512([]byte(password))
	shax := fmt.Sprintf("%x", sha)
	if p != shax {
		return false
	}
	return true
}
