package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	"github.com/jakubc-projects/ustron-work/server/work"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	ClientId, ClientSecret, Issuer, Host string
	SessionService                       work.SessionService
}

type Auth struct {
	config               Config
	oauthConfig          *oauth2.Config
	logoutEndpoint       string
	clientCredentialsCfg *clientcredentials.Config
	idTokenVerifier      *oidc.IDTokenVerifier
	sessionService       work.SessionService
}

func New(cfg Config) *Auth {
	provider, err := oidc.NewProvider(context.Background(), cfg.Issuer)
	if err != nil {
		panic(err)
	}

	endpoint := provider.Endpoint()
	endpoint.AuthStyle = oauth2.AuthStyleInParams

	return &Auth{
		config: cfg,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  fmt.Sprintf("%s/callback", cfg.Host),
			Endpoint:     endpoint,
			Scopes:       []string{"openid", "profile"},
		},
		logoutEndpoint:  fmt.Sprintf("%sv2/logout", cfg.Issuer),
		sessionService:  cfg.SessionService,
		idTokenVerifier: provider.Verifier(&oidc.Config{ClientID: cfg.ClientId, SkipIssuerCheck: true}),
	}
}

func (a Auth) Config() Config {
	return a.config
}

func (a Auth) TokenEndpoint() string {
	return a.oauthConfig.Endpoint.TokenURL
}
