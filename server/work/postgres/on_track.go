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

func (s *OnTrackService) GetOnTrackStatus(ctx context.Context, roundId int) (work.OnTrackStatus, error) {
	result := work.NewOnTrackStatus()

	rows, err := s.db.QueryContext(ctx,
		"SELECT team, status FROM on_track WHERE round_id = $1", roundId)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var entity string
		var status int
		err := rows.Scan(&entity, &status)

		if err != nil {
			return result, err
		}

		result[entity] = status
	}

	return result, nil
}
