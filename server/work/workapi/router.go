package workapi

import "net/http"

func (a *Api) LoadRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/me", a.GetMe)
	mux.HandleFunc("/api/my-registrations", a.MyRegistrations)
}
