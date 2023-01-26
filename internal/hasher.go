package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

func Hasher(path *string) string {
	hasher := sha256.New()
	s, err := os.ReadFile(*path)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
