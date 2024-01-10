package oauthserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/oauth-emulator/pages"
	"github.com/jakubc-projects/ustron-work/oauth-emulator/schema"
	"github.com/lestrrat-go/jwx/jwk"
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

	var personUid uuid.UUID
	if autologin != "" {
		personUid = l.up.RandomUser().Uid
	} else {
		personUidStr := req.PostFormValue("person_uid")
		uid, err := uuid.Parse(personUidStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		personUid = uid
	}

	redirectUriString := req.URL.Query().Get("redirect_uri")
	state := req.URL.Query().Get("state")

	redirectUri, err := url.Parse(redirectUriString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := redirectUri.Query()
	query.Add("code", personUid.String()+"."+randomString(8))
	query.Add("state", state)

	redirectUri.RawQuery = query.Encode()

	http.Redirect(w, req, redirectUri.String(), http.StatusFound)
}

type tokenRequest struct {
	Code     string `schema:"code"`
	ClientId string `schema:"client_id"`
}

func (s *Server) token(w http.ResponseWriter, req *http.Request) {
	time.Sleep(300 * time.Millisecond)

	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body tokenRequest

	if err := schema.Decode(&body, req.Form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	personUidStr, _, _ := strings.Cut(body.Code, ".")

	personUid, err := uuid.Parse(personUidStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := s.newToken(TokenPayload{
		Issuer:    s.serverUrl + "/",
		Audience:  body.ClientId,
		Subject:   personUid.String(),
		PersonUid: personUid,
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
		AuthorizationEndpoint: fmt.Sprintf("%s/authorize", s.serverUrl),
		TokenEndpoint:         fmt.Sprintf("%s/oauth/token", s.serverUrl),
		JwksUri:               fmt.Sprintf("%s/.well-known/jwks.json", s.serverUrl),
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
