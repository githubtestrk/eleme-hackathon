package service

import (
	"cache"
	"crypto/rand"
	"fmt"
	"mem"
)

func newToken() string {
	b := make([]byte, 6)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func CheckTokenLocal(token string) (int, bool) {
	id, ok := mem.GetToken(token)
	if ok {
		return id, true
	} else {
		return 0, false
	}
}

func CheckToken(token string) (int, bool) {
	id, ok := cache.GetToken(token)
	if !ok {
		return 0, false
	} else {
		return id, true
	}
}
