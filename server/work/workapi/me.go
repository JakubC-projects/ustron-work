package workapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/samber/lo"
)

func (a *Api) getMe(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

	me, err := a.personService.GetPerson(ctx, s.PersonID)
	if errors.Is(err, work.ErrNotFound) {
		http.Error(w, "cannot find person", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get person: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(me)
}
