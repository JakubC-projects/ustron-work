package workapi

import (
	"github.com/jakubc-projects/ustron-work/server/work"
)

type Api struct {
	personService       work.PersonService
	registrationService work.RegistrationService
}

func NewApi(ps work.PersonService, rs work.RegistrationService) *Api {
	return &Api{ps, rs}
}
