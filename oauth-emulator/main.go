package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
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
		users.User{DisplayName: "Philly Daly", Uid: uuid.MustParse("15730cac-3f36-4031-bf1e-1d6eedcc4aa0")},
		users.User{DisplayName: "Jakey Boy", Uid: uuid.MustParse("231bce78-4639-477e-a58d-b6a0a2f5d19b")},
		users.User{DisplayName: "Stevie Mallai", Uid: uuid.MustParse("0395e27c-792a-407d-a381-c323c3e61a7a")},
	))

	srv.RegisterHandlers(mux)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func getEnvDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
