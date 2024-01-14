package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type idTokenClaims struct {
	PersonID int `json:"https://login.bcc.no/claims/personId"`
}

const sessionCookieName = "login_session"

func (a *Auth) getSession(req *http.Request) (work.Session, error) {
	cookie, err := req.Cookie(sessionCookieName)
	if err != nil {
		return work.Session{}, fmt.Errorf("no cookie")
	}

	sessionUid, err := uuid.Parse(cookie.Value)
	if err != nil {
		return work.Session{}, fmt.Errorf("invalid session id: %w", err)
	}

	session, err := a.sessionService.GetSession(req.Context(), sessionUid)

	if err != nil {
		return session, fmt.Errorf("cannot find session: %w", err)
	}

	if !session.Expiry.After(time.Now()) {
		return session, fmt.Errorf("session expired")
	}

	return session, nil
}

func (a *Auth) setSession(ctx context.Context, w http.ResponseWriter, session work.Session) error {
	err := a.sessionService.SaveSession(ctx, session)
	if err != nil {
		return fmt.Errorf("cannot save session: %w", err)
	}

	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    session.Uid.String(),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	return nil
}

func (a *Auth) deleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

func getSessionFromIdToken(idToken *oidc.IDToken) (work.Session, error) {
	var claims idTokenClaims

	err := idToken.Claims(&claims)
	if err != nil {
		return work.Session{}, fmt.Errorf("cannot parse claims: %w", err)
	}

	return work.Session{
		Uid:      uuid.New(),
		PersonID: claims.PersonID,
		Expiry:   idToken.Expiry,
	}, nil
}
