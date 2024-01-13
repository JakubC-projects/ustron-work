package auth

import (
	"context"

	"golang.org/x/oauth2"
)

var _ oauth2.TokenSource = &Auth{}

func (a *Auth) Token() (*oauth2.Token, error) {
	return a.clientCredentialsCfg.Token(context.Background())
}
