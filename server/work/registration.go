package work

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work/date"
)

type Registration struct {
	Uid      uuid.UUID `json:"uid"`
	PersonID int       `json:"personID"`
	Team     Team      `json:"team"`

	Type RegistrationType `json:"type"`
	Date date.Date        `json:"date"`

	HourlyWage int     `json:"hourlyWage"`
	Hours      float32 `json:"hours"`

	PaidSum float32 `json:"paidSum"`

	Goal        RegistrationGoal `json:"goal"`
	Description string           `json:"description"`

	CreatedAt time.Time
}

type RegistrationType string

const (
	RegistrationTypeMoney RegistrationType = "Money"
	RegistrationTypeWork  RegistrationType = "Work"
)

type RegistrationGoal string

const (
	RegistrationGoalBuk     RegistrationType = "BUK"
	RegistrationGoalSamvirk RegistrationType = "Samvirk"
)

func (r Registration) Value() float32 {
	return float32(r.HourlyWage)*r.Hours + float32(r.PaidSum)
}

type RegistrationService interface {
	GetPersonRegistrations(context.Context, int) ([]Registration, error)
	GetRegistration(context.Context, uuid.UUID) (Registration, error)
	CreateRegistration(context.Context, Registration) error
	GetStatus(context.Context) (Status, error)
}
