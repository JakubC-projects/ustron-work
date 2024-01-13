package api

import (
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/jakubc-projects/ustron-work/server/work/auth"
)

type Api struct {
	auth          *auth.Auth
	personService work.PersonService
}

func NewApi(auth *auth.Auth, ps work.PersonService) *Api {
	return &Api{auth, ps}
}
