package workapi

import (
	"log/slog"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type Api struct {
	personService       work.PersonService
	registrationService work.RegistrationService
	onTrackService      work.OnTrackService
	roundService        work.RoundService
	logger              *slog.Logger
}

func NewApi(ps work.PersonService, rs work.RegistrationService, ts work.OnTrackService, ros work.RoundService, logger *slog.Logger) *Api {
	return &Api{ps, rs, ts, ros, logger}
}
