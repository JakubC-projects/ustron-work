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

	HourlyWage int `json:"hourlyWage"`
	Hours      int `json:"hours"`

	PaidSum int `json:"paidSum"`

	Comment string `json:"comment"`
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

type Status map[Team]int

func NewStatus() Status {
	return Status{
		TeamBlue:   0,
		TeamGreen:  0,
		TeamOrange: 0,
		TeamRed:    0,
	}
}

type RegistrationService interface {
	GetPersonRegistrations(context.Context, int) ([]Registration, error)
	GetRegistration(context.Context, uuid.UUID) (Registration, error)
	CreateRegistration(context.Context, Registration) error
	UpdateRegistration(context.Context, Registration) error
	GetStatus(context.Context) (Status, error)
}
