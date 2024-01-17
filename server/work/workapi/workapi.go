package workapi

import (
	"github.com/jakubc-projects/ustron-work/server/work"
)

type Api struct {
	personService       work.PersonService
	registrationService work.RegistrationService
	onTrackService      work.OnTrackService
}

func NewApi(ps work.PersonService, rs work.RegistrationService, ts work.OnTrackService) *Api {
	return &Api{ps, rs, ts}
}
