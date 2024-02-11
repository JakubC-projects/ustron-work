package workapi

import (
	"encoding/json"
	"net/http"
)

func (a *Api) status(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.registrationService.GetStatus(ctx)
	if err != nil {

		a.logger.ErrorContext(ctx, "cannot get work status",
			"error", err)
		http.Error(w, "cannot get work status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
