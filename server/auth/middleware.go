package auth

import (
	"fmt"
	"net/http"
)

func (a *Auth) RequiresAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := a.getSession(req)
		if err != nil {
			http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, req)
	})
}

func (a *Auth) ClaimCheck(verifyFunc func(s Session) error, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		session, err := a.getSession(req)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "cannot find valid session")
			return
		}

		if err := verifyFunc(session); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "session invalid: %s", err)
			return
		}

		next.ServeHTTP(w, req)
	})
}
