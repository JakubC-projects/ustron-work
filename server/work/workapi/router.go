package workapi

import "net/http"

func (a *Api) LoadRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/me", a.getMe)
	mux.HandleFunc("GET /api/my-registrations", a.getMyRegistrations)
	mux.HandleFunc("POST /api/my-registrations", a.createMyRegistration)
	mux.HandleFunc("GET /api/status", a.status)
	mux.HandleFunc("GET /api/on-track", a.onTrack)
	mux.HandleFunc("GET /api/on-track-gender", a.onTrackGender)
}
