package auth

import (
	"net/http"

	"github.com/jakubc-projects/ustron-work/server/work"
)

func (a *Auth) RequiresAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s, err := a.getSession(req)
		if err != nil {
			http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)
			return
		}

		req = req.WithContext(work.SetSession(req.Context(), s))

		next.ServeHTTP(w, req)
	})
}
