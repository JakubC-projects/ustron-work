package work

import (
	"context"

	"github.com/google/uuid"
)

type Person struct {
	Uid         uuid.UUID `json:"uid"`
	DisplayName string    `json:"displayName"`
	Team        Team      `json:"team"`
	Role        Role      `json:"role"`
}

type PersonService interface {
	GetPerson(context.Context, uuid.UUID) (Person, error)
	CreatePerson(context.Context, Person) error
	UpdatePerson(context.Context, Person) error
}
