package oauthserver

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/jakubc-projects/ustron-work/oauth-emulator/users"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

type Server struct {
	signingKey jwk.Key
	serverUrl  string
	up         *users.UserProvider
}

func New(signingKey string, serverUrl string, up *users.UserProvider) *Server {
	var jwKey jwk.Key
	var err error

	if signingKey == "" {
		jwKey, err = generateKey()

	} else {
		jwKey, err = jwk.ParseKey([]byte(signingKey))
	}
	if err != nil {
		panic("cannot generate jwk: " + err.Error())
	}

	return &Server{signingKey: jwKey, serverUrl: serverUrl, up: up}
}

func generateKey() (jwk.Key, error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("cannot generate rsa key: %w", err)
	}
	key, err := jwk.New(rsaKey)
	if err != nil {
		return nil, fmt.Errorf("cannot generate jwk key: %w", err)
	}
	jwk.AssignKeyID(key)
	key.Set("alg", jwa.RS256)

	return key, nil
}
