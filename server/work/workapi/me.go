package workapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/samber/lo"
)

func (a *Api) getMe(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

	me, err := a.personService.GetPerson(ctx, s.PersonID)
	if errors.Is(err, work.ErrNotFound) {
		a.logger.WarnContext(ctx, "cannot fetch person",
			"personId", s.PersonID,
			"error", err)
		http.Error(w, "cannot find person", http.StatusNotFound)
		return
	}
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot get person",
			"personId", s.PersonID,
			"error", err)
		http.Error(w, "error fetching person", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(me)
}
