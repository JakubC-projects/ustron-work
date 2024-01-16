package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func (a *Auth) SetLoginHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/login", a.loginHandler)
	mux.HandleFunc("/callback", a.callbackHandler)
	mux.HandleFunc("/logout", a.logoutHandler)
}

func (a *Auth) loginHandler(w http.ResponseWriter, req *http.Request) {
	state, err := generateRandomState()
	if err != nil {
		errCtx := fmt.Errorf("cannot generate random state: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errCtx.Error())
		return
	}

	cookie := http.Cookie{Name: "state", Value: state, SameSite: http.SameSiteLaxMode, Expires: time.Now().AddDate(0, 0, 1)}
	http.SetCookie(w, &cookie)

	http.Redirect(w, req, a.oauthConfig.AuthCodeURL(state, oauth2.SetAuthURLParam("response_mode", "post_form")), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func (a *Auth) callbackHandler(w http.ResponseWriter, req *http.Request) {
	token, err := a.getCallbackToken(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "cannot get callback token: %s", err)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "missing id token")
		return
	}

	idToken, err := a.idTokenVerifier.Verify(req.Context(), rawIDToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "invalid id token: %s", err)
		return
	}

	session, err := getSessionFromIdToken(idToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "cannot create session from token: %s", err)
		return
	}

	if err := a.setSession(req.Context(), w, session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "cannot save session: %s", err)
		return
	}

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func (a *Auth) getCallbackToken(req *http.Request) (*oauth2.Token, error) {
	state, err := req.Cookie("state")
	if err != nil {
		return nil, fmt.Errorf("invalid state parameter: %w", err)
	}

	if req.URL.Query().Get("state") != state.Value {
		return nil, fmt.Errorf("invalid state parameter")
	}

	code := req.URL.Query().Get("code")

	token, err := a.oauthConfig.Exchange(req.Context(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange an authorization code for a token: %w", err)
	}

	return token, nil
}

func (a *Auth) logoutHandler(w http.ResponseWriter, req *http.Request) {
	a.deleteSession(w)
	http.Redirect(w, req, a.logoutEndpoint, http.StatusTemporaryRedirect)
}
