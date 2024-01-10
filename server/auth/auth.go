package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/securecookie"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	ClientId, ClientSecret, Issuer, Host string
}

type Auth struct {
	config               Config
	oauthConfig          *oauth2.Config
	logoutEndpoint       string
	clientCredentialsCfg *clientcredentials.Config
	cookieEncryption     *securecookie.SecureCookie
	idTokenVerifier      *oidc.IDTokenVerifier
}

func New(cfg Config) *Auth {
	provider, err := oidc.NewProvider(context.Background(), cfg.Issuer)
	if err != nil {
		panic(err)
	}

	return &Auth{
		config: cfg,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  fmt.Sprintf("%s/callback", cfg.Host),
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{"openid", "profile", "roles"},
		},
		logoutEndpoint:   fmt.Sprintf("%sv2/logout", cfg.Issuer),
		cookieEncryption: securecookie.New([]byte(cfg.ClientSecret), nil).SetSerializer(securecookie.JSONEncoder{}),
		idTokenVerifier:  provider.Verifier(&oidc.Config{ClientID: cfg.ClientId, SkipIssuerCheck: true}),
	}
}

func (a Auth) Config() Config {
	return a.config
}

func (a Auth) TokenEndpoint() string {
	return a.oauthConfig.Endpoint.TokenURL
}
