package work

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Uid      uuid.UUID
	PersonID int
	Expiry   time.Time
}

type SessionService interface {
	GetSession(context.Context, uuid.UUID) (Session, error)
	SaveSession(context.Context, Session) error
}

type sessionKey struct{}

func SetSession(ctx context.Context, s Session) context.Context {
	return context.WithValue(ctx, sessionKey{}, s)
}

func GetSession(ctx context.Context) (Session, bool) {
	v := ctx.Value(sessionKey{})
	session, ok := v.(Session)
	return session, ok
}
