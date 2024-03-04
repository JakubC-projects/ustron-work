package workapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/samber/lo"
)

func (a *Api) createMyRegistration(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

	var registration work.Registration

	p, err := a.personService.GetPerson(ctx, s.PersonID)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot fetch person",
			"personId", s.PersonID,
			"error", err)
		http.Error(w, "cannot fetch person", http.StatusInternalServerError)
		return

	}

	if err := json.NewDecoder(req.Body).Decode(&registration); err != nil {
		err := fmt.Errorf("cannot decode registration body: %w", err)
		a.logger.WarnContext(ctx, "invalid request json",
			"error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	registration.Uid = uuid.New()
	registration.PersonID = s.PersonID
	registration.Team = p.Team
	registration.CreatedAt = time.Now()

	if err := validateCreatedRegistration(&registration); err != nil {
		err := fmt.Errorf("validation failed: %w", err)
		a.logger.WarnContext(ctx, "registration validation failed",
			"registration", registration,
			"error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	if err := a.registrationService.CreateRegistration(ctx, registration); err != nil {
		a.logger.ErrorContext(ctx, "cannot save registration",
			"registration", registration,
			"error", err)
		http.Error(w, "cannot save registration", http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(registration)
}

func (a *Api) getMyRegistrations(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	s := lo.Must(work.GetSession(ctx))

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

	registrations, err := a.registrationService.GetPersonRegistrations(ctx, s.PersonID, round)
	if err != nil {
		a.logger.ErrorContext(ctx, "cannot fetch registrations",
			"personId", s.PersonID,
			"error", err,
		)
		http.Error(w, "cannot fetch registrations", http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(registrations)
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
