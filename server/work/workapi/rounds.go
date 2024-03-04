package workapi

import (
	"encoding/json"
	"net/http"
)

func (a *Api) getRounds(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	rounds, err := a.roundService.GetRounds(ctx)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot get rounds",
			"error", err)
		http.Error(w, "cannot get rounds", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rounds)
}
