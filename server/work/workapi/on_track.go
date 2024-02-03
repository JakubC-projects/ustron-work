package workapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *Api) onTrack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.onTrackService.GetOnTrackStatus(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get on track status: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (a *Api) onTrackGender(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.onTrackService.GetOnTrackGenderStatus(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get on track status: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
