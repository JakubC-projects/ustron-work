package api

import "net/http"

func (a *Api) LoadRoutes(mux *http.ServeMux) {
	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("/api/me", a.GetMeRoute)

	mux.Handle("/api/", a.auth.RequiresAuth(apiRouter))
}
