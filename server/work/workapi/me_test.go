package workapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/jakubc-projects/ustron-work/server/work/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	person := work.Person{
		PersonID:    54512,
		DisplayName: "Philly",
		Team:        work.TeamBlue,
	}
	ps := mock.NewPersonService(person)

	api := NewApi(ps, nil, nil)

	t.Run("OK", func(t *testing.T) {
		loggedInSession := work.SetSession(context.Background(), work.Session{
			Uid:      uuid.Nil,
			PersonID: person.PersonID,
		})
		req := httptest.NewRequest(http.MethodGet, "/api/me", nil).WithContext(loggedInSession)
		w := httptest.NewRecorder()
		api.getMe(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res work.Person
		json.NewDecoder(w.Body).Decode(&res)
		assert.Equal(t, person, res)
	})

	t.Run("Not Found", func(t *testing.T) {
		loggedInSession := work.SetSession(context.Background(), work.Session{
			Uid:      uuid.Nil,
			PersonID: 1,
		})
		req := httptest.NewRequest(http.MethodGet, "/api/me", nil).WithContext(loggedInSession)
		w := httptest.NewRecorder()
		api.getMe(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})
}
