package workapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/samber/lo"
)

func (a *Api) myRegistrations(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

	var res any
	var err error

	switch req.Method {
	case http.MethodPost:
		res, err = a.createMyRegistration(ctx, s, req)
	case http.MethodGet:
		res, err = a.getMyRegistrations(ctx, s, req)

	default:
		err = errors.New("method not allowed")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res != nil {
		json.NewEncoder(w).Encode(res)
	}
}

func (a *Api) createMyRegistration(ctx context.Context, s work.Session, req *http.Request) (work.Registration, error) {
	var registration work.Registration

	p, err := a.personService.GetPerson(ctx, s.PersonID)
	if err != nil {
		return registration, fmt.Errorf("cannot get person: %w", err)
	}

	if err := json.NewDecoder(req.Body).Decode(&registration); err != nil {
		return registration, fmt.Errorf("cannot decode registration body: %w", err)
	}

	registration.Uid = uuid.New()
	registration.PersonID = s.PersonID
	registration.Team = p.Team
	registration.CreatedAt = time.Now()

	if err := validateCreatedRegistration(&registration); err != nil {
		return registration, fmt.Errorf("validation failed: %w", err)
	}

	if err := a.registrationService.CreateRegistration(ctx, registration); err != nil {
		return registration, fmt.Errorf("cannot save registration: %w", err)
	}

	return registration, nil
}

func (a *Api) getMyRegistrations(ctx context.Context, s work.Session, req *http.Request) ([]work.Registration, error) {

	registrations, err := a.registrationService.GetPersonRegistrations(ctx, s.PersonID)
	if err != nil {
		return registrations, fmt.Errorf("cannot get my registrations: %w", err)
	}

	return registrations, nil
}

func validateCreatedRegistration(reg *work.Registration) error {
	if reg.Date.After(civil.DateOf(time.Now())) {
		return fmt.Errorf("registration date cannot be in the future")
	}

	if reg.Type == work.RegistrationTypeWork {
		if reg.HourlyWage <= 0 {
			return fmt.Errorf("hourly wage has to be greater than 0")
		}
		if reg.Hours <= 0 {
			return fmt.Errorf("hours cannot be negative")
		}
		if reg.HourlyWage > 200 {
			return fmt.Errorf("too high hourly wage, nobody makes that much 🧐")
		}
		if reg.Hours > 24 {
			return fmt.Errorf("too many hours, nobody works that much. Now go get some sleep 😴")
		}
		if len(reg.Description) == 0 {
			return fmt.Errorf("please provide work description")
		}

		reg.PaidSum = 0
	} else {
		if reg.PaidSum <= 0 {
			return fmt.Errorf("paid amount has to be greater than 0")
		}
		if reg.PaidSum > 1000000 {
			return fmt.Errorf("too high amount, nobody is that rich 🤑")
		}

		reg.HourlyWage = 0
		reg.Hours = 0
	}

	return nil
}
