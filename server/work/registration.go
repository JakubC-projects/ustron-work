package work

import (
	"context"

	"github.com/google/uuid"
)

type Registration struct {
	Uid              uuid.UUID `json:"uid"`
	PersonUid        uuid.UUID `json:"personUid"`
	Team             Team      `json:"team"`
	RegistrationType RegistrationType

	HourlyWage int `json:"hourlyWage"`
	Hours      int `json:"hours"`

	PaidSum int `json:"paidSum"`
}

type RegistrationType string

const (
	RegistrationTypeMoney RegistrationType = "Money"
	RegistrationTypeWork  RegistrationType = "Work"
)

func (r Registration) Value() int {
	switch r.RegistrationType {
	case RegistrationTypeWork:
		return r.HourlyWage * r.Hours
	case RegistrationTypeMoney:
		return r.PaidSum
	}
	return 0
}

type RegistrationService interface {
	GetPersonRegistrations(context.Context, uuid.UUID) ([]Registration, error)
	GetRegistration(context.Context, uuid.UUID) (Registration, error)
	CreateRegistration(context.Context, Registration) error
	UpdateRegistration(context.Context, Registration) error
}
