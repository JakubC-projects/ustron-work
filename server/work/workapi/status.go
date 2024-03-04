package workapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *Api) status(w http.ResponseWriter, req *http.Request) {
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

	round, err := a.roundService.GetRound(ctx, roundId)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot find round",
			"roundId", roundIdStr,
			"error", err,
		)
		http.Error(w, "cannot find round", http.StatusBadRequest)
		return
	}

	status, err := a.registrationService.GetStatus(ctx, round)
	if err != nil {

		a.logger.ErrorContext(ctx, "cannot get work status",
			"error", err)
		http.Error(w, "cannot get work status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
