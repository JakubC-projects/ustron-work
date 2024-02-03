package workapi

import (
	"bytes"
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

func TestMyRegistrations(t *testing.T) {
	person := work.Person{
		PersonID:    54512,
		DisplayName: "Philly",
		Team:        work.TeamBlue,
	}

	registration := work.Registration{
		Uid:        uuid.New(),
		PersonID:   person.PersonID,
		Team:       person.Team,
		Type:       work.RegistrationTypeWork,
		HourlyWage: 20,
		Hours:      3,
	}

	ps := mock.NewPersonService(person)
	rs := mock.NewRegistrationService(registration)

	api := NewApi(ps, rs, nil)

	loggedInSession := work.SetSession(context.Background(), work.Session{
		Uid:      uuid.Nil,
		PersonID: person.PersonID,
	})

	t.Run("Find OK", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/my-registrations", nil).WithContext(loggedInSession)
		w := httptest.NewRecorder()

		api.myRegistrations(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res []work.Registration
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, []work.Registration{registration}, res)
	})

	t.Run("Create Ok", func(t *testing.T) {
		registration := work.Registration{
			Type:    work.RegistrationTypeMoney,
			PaidSum: 100,
		}

		b := &bytes.Buffer{}
		json.NewEncoder(b).Encode(registration)

		req := httptest.NewRequest(http.MethodPost, "/api/my-registrations", b).WithContext(loggedInSession)
		w := httptest.NewRecorder()
		api.myRegistrations(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res work.Registration
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, registration.Type, res.Type)
		assert.Equal(t, registration.PaidSum, res.PaidSum)
		assert.Equal(t, person.PersonID, res.PersonID)
		assert.Equal(t, person.Team, res.Team)

		assert.Len(t, rs.S.Data, 2)
		assert.Equal(t, res.Uid, rs.S.Data[1].Uid)
	})

}
