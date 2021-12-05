package util

import (
	"crypto/rand"
	"encoding/base64"
)

// TODO: @hammer, 걍 아무거나 긁어옴 위치 옮기고 개선해야함

type IDGenerator interface {
	NewID() (string, error)
}

type DefaultGenerator struct {
}

func (g *DefaultGenerator) NewID() (string, error) {
	s, err := generateRandomStringURLSafe(10)
	if err != nil {
		return "", err
	}
	return s, nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomStringURLSafe(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
