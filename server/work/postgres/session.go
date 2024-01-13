package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type SessionService struct {
	db *sql.DB
}

func NewSessionService(db *sql.DB) *SessionService {
	return &SessionService{db}
}

var _ work.SessionService = (*SessionService)(nil)

func (srv *SessionService) GetSession(ctx context.Context, uid uuid.UUID) (work.Session, error) {
	var p work.Session

	rows, err := srv.db.QueryContext(ctx, "SELECT uid, person_uid, expiry FROM Sessions WHERE uid = $1", uid)

	if err != nil {
		return p, fmt.Errorf("sql error getting Session: %w", err)
	}

	sessions, err := scanSessions(rows)
	if err != nil {
		return work.Session{}, fmt.Errorf("error scanning registrations: %w", err)
	}

	if len(sessions) != 1 {
		return work.Session{}, work.ErrNotFound
	}

	return sessions[0], nil
}

func (srv *SessionService) SaveSession(ctx context.Context, s work.Session) error {
	_, err := srv.db.ExecContext(ctx, "INSERT INTO Sessions (uid, person_uid, expiry) VALUES ($1, $2, $3)", s.Uid, s.PersonUid, s.Expiry)

	return err
}

func scanSessions(rows *sql.Rows) ([]work.Session, error) {

	var sessions []work.Session

	for rows.Next() {
		var p work.Session

		err := rows.Scan(&p.Uid, &p.PersonUid, &p.Expiry)

		if err != nil {
			return nil, err
		}
		sessions = append(sessions, p)
	}
	return sessions, nil
}
