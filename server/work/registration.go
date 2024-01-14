package work

import (
	"context"

	"github.com/google/uuid"
)

type Registration struct {
	Uid      uuid.UUID        `json:"uid"`
	PersonID int              `json:"personID"`
	Team     Team             `json:"team"`
	Type     RegistrationType `json:"type"`

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
	switch r.Type {
	case RegistrationTypeWork:
		return r.HourlyWage * r.Hours
	case RegistrationTypeMoney:
		return r.PaidSum
	}
	return 0
}

type RegistrationService interface {
	GetPersonRegistrations(context.Context, int) ([]Registration, error)
	GetRegistration(context.Context, uuid.UUID) (Registration, error)
	CreateRegistration(context.Context, Registration) error
	UpdateRegistration(context.Context, Registration) error
}
