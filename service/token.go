package service

import (
	"cache"
	"crypto/rand"
	"fmt"
)

func newToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func CheckToken(token string) (int, bool) {
	id, ok := cache.GetToken(token)
	if !ok {
		return 0, false
	} else {
		return id, true
	}
}
