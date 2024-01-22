package workapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *Api) status(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.registrationService.GetStatus(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get my registrations: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
