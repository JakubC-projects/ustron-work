package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/google/uuid"
)

type Session struct {
	PersonUid uuid.UUID
	Roles     []Role
	ExpiresAt time.Time
}

type idTokenClaims struct {
	PersonUid uuid.UUID `json:"https://login.bcc.no/claims/personUid"`
}

type Role struct {
	OrgId    int       `json:"orgId"`
	OrgUid   uuid.UUID `json:"orgUid"`
	RoleName string    `json:"roleName"`
}

const sessionCookieName = "login_session"

func (a *Auth) getSession(req *http.Request) (Session, error) {
	var s Session
	cookie, err := req.Cookie(sessionCookieName)
	if err != nil {
		return s, fmt.Errorf("no cookie")
	}

	err = a.cookieEncryption.Decode(sessionCookieName, cookie.Value, &s)

	if err != nil {
		return s, fmt.Errorf("cannot decode session: %w", err)

	}

	if !s.ExpiresAt.After(time.Now()) {
		return s, fmt.Errorf("session expired")
	}

	return s, nil
}

func (a *Auth) setSession(w http.ResponseWriter, session Session) error {
	encoded, err := a.cookieEncryption.Encode(sessionCookieName, session)
	if err != nil {
		return fmt.Errorf("cannot encode cookie value:%w", err)
	}
	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    encoded,
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

func getSessionFromIdToken(idToken *oidc.IDToken) (Session, error) {
	var claims idTokenClaims

	err := idToken.Claims(&claims)
	if err != nil {
		return Session{}, fmt.Errorf("cannot parse claims: %w", err)
	}

	return Session{
		ExpiresAt: idToken.Expiry,
		PersonUid: claims.PersonUid,
		Roles:     []Role{},
	}, nil
}
