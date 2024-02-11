package workapi

import (
	"log/slog"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type Api struct {
	personService       work.PersonService
	registrationService work.RegistrationService
	onTrackService      work.OnTrackService
	logger              *slog.Logger
}

func NewApi(ps work.PersonService, rs work.RegistrationService, ts work.OnTrackService, logger *slog.Logger) *Api {
	return &Api{ps, rs, ts, logger}
}
