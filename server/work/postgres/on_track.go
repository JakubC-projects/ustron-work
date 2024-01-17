package postgres

import (
	"context"
	"database/sql"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type OnTrackService struct {
	db *sql.DB
}

func NewOnTrackService(db *sql.DB) *OnTrackService {
	return &OnTrackService{db}
}

var _ work.OnTrackService = (*OnTrackService)(nil)

func (s *OnTrackService) GetOnTrackStatus(ctx context.Context) (work.Status, error) {
	result := work.NewStatus()

	rows, err := s.db.QueryContext(ctx, "SELECT team, status FROM on_track")

	if err != nil {
		return result, err
	}
	for rows.Next() {
		var team work.Team
		var status int
		err := rows.Scan(&team, &status)

		if err != nil {
			return result, err
		}

		result[team] = status
	}

	return result, nil
}
