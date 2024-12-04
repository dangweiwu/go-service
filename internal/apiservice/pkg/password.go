package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func GetPassword(x, salt string) (string, error) {
	hasher := sha256.New()
	if _, err := io.WriteString(hasher, x); err != nil {
		return "", err
	}

	if _, err := io.WriteString(hasher, salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
