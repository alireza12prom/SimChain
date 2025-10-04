package utility

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(text string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
