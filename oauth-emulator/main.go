package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jakubc-projects/ustron-work/oauth-emulator/oauthserver"
	"github.com/jakubc-projects/ustron-work/oauth-emulator/users"
)

var (
	port       = os.Getenv("PORT")
	serverHost = os.Getenv("SERVER_HOST")
	signingKey = os.Getenv("SIGNING_KEY")
)

func main() {
	mux := http.NewServeMux()
	srv := oauthserver.New(signingKey, serverHost, users.NewUserProvider(
		users.User{DisplayName: "Test User Green", PersonID: 54512},
		users.User{DisplayName: "Test User Red", PersonID: 54513},
		users.User{DisplayName: "Test User Blue", PersonID: 54514},
	))

	srv.RegisterHandlers(mux)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
