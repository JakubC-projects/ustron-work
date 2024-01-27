package workapi

import "net/http"

func (a *Api) LoadRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/me", a.getMe)
	mux.HandleFunc("/api/my-registrations", a.myRegistrations)
	mux.HandleFunc("/api/status", a.status)
	mux.HandleFunc("/api/on-track", a.onTrack)
	mux.HandleFunc("/api/on-track-gender", a.onTrackGender)
}
