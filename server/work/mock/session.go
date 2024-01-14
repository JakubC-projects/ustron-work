package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type SessionService struct {
	S DataService[work.Session, uuid.UUID]
}

var _ work.SessionService = (*SessionService)(nil)

func NewSessionService(data ...work.Session) *SessionService {
	return &SessionService{
		S: NewDataService(data, func(d work.Session) uuid.UUID { return d.Uid }),
	}
}

func (rs *SessionService) GetSession(ctx context.Context, id uuid.UUID) (work.Session, error) {
	return rs.S.Get(ctx, id)
}
func (rs *SessionService) SaveSession(ctx context.Context, reg work.Session) error {
	return rs.S.Create(ctx, reg)
}
