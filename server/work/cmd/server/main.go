package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jakubc-projects/ustron-work/server/work/api"
	"github.com/jakubc-projects/ustron-work/server/work/auth"
	"github.com/jakubc-projects/ustron-work/server/work/frontend"
	"github.com/jakubc-projects/ustron-work/server/work/postgres"
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
	connectionString  = os.Getenv("POSTGERS_CONNECTIONSTRING")
)

func main() {

	db, err := postgres.NewDb(connectionString)
	if err != nil {
		panic(fmt.Errorf("cannot open db connection: %w", err))
	}

	ps := postgres.NewPersonService(db)
	ss := postgres.NewSessionService(db)
	// rs := postgres.NewRegistrationService(db)

	auth := auth.New(
		auth.Config{
			ClientId:       oauthClientId,
			ClientSecret:   oauthClientSecret,
			Issuer:         oauthIssuer,
			Host:           serverHost,
			SessionService: ss,
		},
	)

	api := api.NewApi(auth, ps)

	mux := http.NewServeMux()

	auth.SetLoginHandlers(mux)
	api.LoadRoutes(mux)

	authenticatedMux := http.NewServeMux()
	mux.Handle("/", auth.RequiresAuth(authenticatedMux))

	authenticatedMux.Handle("/", (frontend.Handler(frontendLocation)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
