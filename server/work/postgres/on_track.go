package postgres

import (
	"context"
	"database/sql"
	"fmt"

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

func (s *OnTrackService) SetOnTrackStatus(ctx context.Context, status work.Status) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE on_track"); err != nil {
		return fmt.Errorf("cannot remove existing status: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO on_track (team, status) VALUES 
		('Blue', $1), ('Green', $2), ('Orange', $3), ('Red', $4)`,
		status[work.TeamBlue], status[work.TeamGreen], status[work.TeamOrange], status[work.TeamRed])
	if err != nil {
		return fmt.Errorf("cannot create new status: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit transaction: %w", err)
	}

	return nil
}
