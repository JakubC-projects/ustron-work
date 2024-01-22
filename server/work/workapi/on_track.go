package workapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/samber/lo"
)

func (a *Api) onTrack(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		a.getOnTrack(w, req)
	case http.MethodPost:
		a.setOnTrack(w, req)
	}
}

func (a *Api) getOnTrack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	status, err := a.onTrackService.GetOnTrackStatus(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get on track status: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (a *Api) setOnTrack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

	person, err := a.personService.GetPerson(ctx, s.PersonID)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot find person: %s", err), http.StatusInternalServerError)
		return
	}

	if person.Role != work.RoleAdmin {
		http.Error(w, "missing permissions to update on track status", http.StatusForbidden)
		return
	}

	var status work.Status
	err = json.NewDecoder(req.Body).Decode(&status)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid on track status: %s", err), http.StatusBadRequest)
		return
	}

	err = a.onTrackService.SetOnTrackStatus(ctx, status)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot update on track status: %s", err), http.StatusInternalServerError)
		return
	}
}
