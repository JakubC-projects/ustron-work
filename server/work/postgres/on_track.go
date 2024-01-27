package postgres

import (
	"context"
	"database/sql"
	"fmt"

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

func (s *OnTrackService) SetOnTrackStatus(ctx context.Context, status work.Status) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "DELETE FROM on_track WHERE team = ANY ($1)", pq.Array(work.Teams)); err != nil {
		return fmt.Errorf("cannot remove existing status: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO on_track (team, status) VALUES 
		($1, $2), ($3, $4), ($5, $6), ($7, $8)`,
		work.TeamBlue, status[work.TeamBlue],
		work.TeamGreen, status[work.TeamGreen],
		work.TeamOrange, status[work.TeamOrange],
		work.TeamRed, status[work.TeamRed])

	if err != nil {
		return fmt.Errorf("cannot create new status: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit transaction: %w", err)
	}

	return nil
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

func (s *OnTrackService) SetOnTrackGenderStatus(ctx context.Context, status work.GenderStatus) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "DELETE FROM on_track WHERE team = ANY ($1)", pq.Array(work.Genders)); err != nil {
		return fmt.Errorf("cannot remove existing status: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO on_track (team, status) VALUES 
		($1, $2), ($3, $4)`,
		work.GenderMale, status[work.GenderMale],
		work.GenderFemale, status[work.GenderFemale],
	)
	if err != nil {
		return fmt.Errorf("cannot create new status: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit transaction: %w", err)
	}

	return nil
}
