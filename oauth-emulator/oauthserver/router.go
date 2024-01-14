package oauthserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jakubc-projects/ustron-work/oauth-emulator/pages"
	"github.com/jakubc-projects/ustron-work/oauth-emulator/schema"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/samber/lo"
)

func (l *Server) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/authorize", l.authorize)
	mux.HandleFunc("/oauth/token", l.token)
	mux.HandleFunc("/.well-known/openid-configuration", l.configuration)
	mux.HandleFunc("/.well-known/jwks.json", l.jwks)
}

func (l *Server) authorize(w http.ResponseWriter, req *http.Request) {
	autologin := req.URL.Query().Get("autologin")
	if autologin == "" && req.Method == http.MethodGet {
		pages.RenderLogin(w, pages.Login{AvailableUsers: l.up.Users})
		return
	}

	var personID int
	if autologin != "" {
		personID = l.up.RandomUser().PersonID
	} else {
		personIDStr := req.PostFormValue("person_id")

		var err error
		personID, err = strconv.Atoi(personIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	redirectUriString := req.URL.Query().Get("redirect_uri")
	state := req.URL.Query().Get("state")

	redirectUri, err := url.Parse(redirectUriString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := redirectUri.Query()
	query.Add("code", fmt.Sprintf("%d.%s", personID, randomString(8)))
	query.Add("state", state)

	redirectUri.RawQuery = query.Encode()

	http.Redirect(w, req, redirectUri.String(), http.StatusFound)
}

type tokenRequest struct {
	Code     string `schema:"code"`
	ClientId string `schema:"client_id"`
}

func (s *Server) token(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body tokenRequest

	if err := schema.Decode(&body, req.Form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	personIDStr, _, _ := strings.Cut(body.Code, ".")

	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := s.newToken(TokenPayload{
		Issuer:   s.serverUrl,
		Audience: body.ClientId,
		Subject:  strconv.Itoa(personID),
		PersonID: personID,
	}, time.Hour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := TokenResponse{
		AccessToken: token,
		IdToken:     token,
		ExpiresIn:   3000,
		TokenType:   "Bearer",
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (s *Server) configuration(w http.ResponseWriter, req *http.Request) {

	response := OidcDiscovery{
		Issuer:                s.serverUrl,
		AuthorizationEndpoint: lo.Must(url.JoinPath(s.serverUrl, "authorize")),
		TokenEndpoint:         lo.Must(url.JoinPath(s.serverUrl, "oauth", "token")),
		JwksUri:               lo.Must(url.JoinPath(s.serverUrl, ".well-known", "jwks.json")),
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (s *Server) jwks(w http.ResponseWriter, req *http.Request) {
	keySet := jwk.NewSet()
	public, err := s.signingKey.PublicKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	keySet.Add(public)

	json.NewEncoder(w).Encode(keySet)
}
