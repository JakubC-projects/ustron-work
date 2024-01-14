package mock

import (
	"context"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type StatusService struct {
	rs *RegistrationService
}

func NewStatusService(rs *RegistrationService) *StatusService {
	return &StatusService{rs}
}

var _ work.StatusService = (*StatusService)(nil)

func (s *StatusService) GetStatus(context.Context) (work.Status, error) {
	status := work.Status{
		work.TeamBlue:   0,
		work.TeamGreen:  0,
		work.TeamOrange: 0,
		work.TeamRed:    0,
	}

	for _, r := range s.rs.S.Data {
		status[r.Team] += r.Value()
	}

	return status, nil
}
