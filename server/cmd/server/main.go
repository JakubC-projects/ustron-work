package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jakubc-projects/ustron-work/server/auth"
	"github.com/jakubc-projects/ustron-work/server/frontend"
)

type Config struct {
	Server struct {
		Port int    `envconfig:"PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Oauth struct {
		ClientID     string `envconfig:"OAUTH_CLIENT_ID"`
		ClientSecret string `envconfig:"OAUTH_CLIENT_SECRET"`
		Issuer       string `envconfig:"OAUTH_ISSUER"`
	}
	Coreapi struct {
		BasePath string `envconfig:"COREAPI_BASE_PATH"`
		Audience string `envconfig:"COREAPI_AUDIENCE"`
	}
	FrontendLocation string `envconfig:"FRONTEND_LOCATION"`
}

var (
	port              = os.Getenv("PORT")
	serverHost        = os.Getenv("SERVER_HOST")
	oauthClientId     = os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	oauthIssuer       = os.Getenv("OAUTH_ISSUER")
	frontendLocation  = os.Getenv("FRONTEND_LOCATION")
)

func main() {

	a := auth.New(
		auth.Config{
			ClientId:     oauthClientId,
			ClientSecret: oauthClientSecret,
			Issuer:       oauthIssuer,
			Host:         serverHost,
		},
	)

	mux := http.NewServeMux()

	a.SetLoginHandlers(mux)

	authenticatedMux := http.NewServeMux()
	mux.Handle("/", a.RequiresAuth(authenticatedMux))

	authenticatedMux.Handle("/", (frontend.Handler(frontendLocation)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
