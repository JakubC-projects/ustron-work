package workapi

import (
	"encoding/json"
	"net/http"
)

func (a *Api) onTrack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.onTrackService.GetOnTrackStatus(ctx)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot get on track status",
			"error", err)
		http.Error(w, "cannot get on track status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (a *Api) onTrackGender(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.onTrackService.GetOnTrackGenderStatus(ctx)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot get on track gender status",
			"error", err)
		http.Error(w, "cannot get on track gender status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
