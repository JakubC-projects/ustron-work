package work

import (
	"context"

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

	PaidSum int `json:"paidSum"`

	Description string `json:"description"`
}

type RegistrationType string

const (
	RegistrationTypeMoney RegistrationType = "Money"
	RegistrationTypeWork  RegistrationType = "Work"
)

func (r Registration) Value() float32 {
	return float32(r.HourlyWage)*r.Hours + float32(r.PaidSum)
}

type RegistrationService interface {
	GetPersonRegistrations(context.Context, int) ([]Registration, error)
	GetRegistration(context.Context, uuid.UUID) (Registration, error)
	CreateRegistration(context.Context, Registration) error
	UpdateRegistration(context.Context, Registration) error
	GetStatus(context.Context) (Status, error)
}
