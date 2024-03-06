package workapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *Api) onTrack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	roundIdStr := req.URL.Query().Get("roundId")
	roundId, err := strconv.Atoi(roundIdStr)
	if err != nil {
		a.logger.ErrorContext(ctx, "invalid round id",
			"roundId", roundIdStr,
			"error", err,
		)
		http.Error(w, "cannot fetch registrations", http.StatusBadRequest)
		return
	}

	status, err := a.onTrackService.GetOnTrackStatus(ctx, roundId)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot get on track status",
			"error", err)
		http.Error(w, "cannot get on track status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
