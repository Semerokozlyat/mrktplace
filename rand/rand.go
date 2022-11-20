package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(len int) ([]byte, error) {
	b := make([]byte, len)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	if nRead < len {
		return nil, fmt.Errorf("insufficient number of bytes read")
	}
	return b, nil
}

func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("producing random bytes: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

const SessionTokenBytes = 32

func SessionToken() (string, error) {
	return String(SessionTokenBytes)
}
