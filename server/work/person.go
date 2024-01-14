package work

import (
	"context"
)

type Person struct {
	PersonID    int    `json:"personID"`
	DisplayName string `json:"displayName"`
	Team        Team   `json:"team"`
	Role        Role   `json:"role"`
}

type PersonService interface {
	GetPerson(context.Context, int) (Person, error)
	CreatePerson(context.Context, Person) error
	UpdatePerson(context.Context, Person) error
}
