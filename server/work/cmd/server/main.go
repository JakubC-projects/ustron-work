package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jakubc-projects/ustron-work/server/work/auth"
	"github.com/jakubc-projects/ustron-work/server/work/frontend"
	"github.com/jakubc-projects/ustron-work/server/work/postgres"
	"github.com/jakubc-projects/ustron-work/server/work/workapi"
)

var (
	port                = os.Getenv("PORT")
	serverHost          = os.Getenv("SERVER_HOST")
	oauthClientId       = os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret   = os.Getenv("OAUTH_CLIENT_SECRET")
	oauthIssuer         = os.Getenv("OAUTH_ISSUER")
	oauthLogoutEndpoint = os.Getenv("OAUTH_LOGOUT_ENDPOINT")
	frontendLocation    = os.Getenv("FRONTEND_LOCATION")
	connectionString    = os.Getenv("POSTGERS_CONNECTIONSTRING")
)

func main() {

	db, err := postgres.NewDb(connectionString)
	if err != nil {
		panic(fmt.Errorf("cannot open db connection: %w", err))
	}

	ps := postgres.NewPersonService(db)
	ss := postgres.NewSessionService(db)
	rs := postgres.NewRegistrationService(db)
	ts := postgres.NewOnTrackService(db)

	auth := auth.New(
		auth.Config{
			ClientId:       oauthClientId,
			ClientSecret:   oauthClientSecret,
			Issuer:         oauthIssuer,
			LogoutEndpoint: oauthLogoutEndpoint,
			Host:           serverHost,
			SessionService: ss,
		},
	)

	api := workapi.NewApi(ps, rs, ts)

	mux := http.NewServeMux()

	auth.SetLoginHandlers(mux)

	authenticatedMux := http.NewServeMux()

	mux.Handle("/", auth.RequiresAuth(authenticatedMux))

	api.LoadRoutes(authenticatedMux)
	authenticatedMux.Handle("/", (frontend.Handler(frontendLocation)))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
