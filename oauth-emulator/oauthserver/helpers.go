package oauthserver

import (
	"crypto/rand"
	"math/big"
)

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		bigN, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic("cannot generate random number")
		}
		b[i] = letters[bigN.Int64()]
	}
	return string(b)
}
