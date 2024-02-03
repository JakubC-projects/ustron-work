package postgres

import (
	"context"
	"database/sql"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/lib/pq"
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

	rows, err := s.db.QueryContext(ctx, "SELECT team, status FROM on_track WHERE team = ANY ($1)", pq.Array(work.Teams))

	if err != nil {
		return result, err
	}
	for rows.Next() {
		var team work.Team
		var status float32
		err := rows.Scan(&team, &status)

		if err != nil {
			return result, err
		}

		result[team] = status
	}

	return result, nil
}

func (s *OnTrackService) GetOnTrackGenderStatus(ctx context.Context) (work.GenderStatus, error) {
	result := work.NewGenderStatus()

	rows, err := s.db.QueryContext(ctx, "SELECT team, status FROM on_track WHERE team = ANY ($1)", pq.Array(work.Genders))

	if err != nil {
		return result, err
	}
	for rows.Next() {
		var gender work.Gender
		var status float32
		err := rows.Scan(&gender, &status)

		if err != nil {
			return result, err
		}

		result[gender] = status
	}

	return result, nil
}
