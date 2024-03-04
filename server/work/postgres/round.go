package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type RoundService struct {
	db *sql.DB
}

func NewRoundService(db *sql.DB) *RoundService {
	return &RoundService{db}
}

func (s *RoundService) GetRounds(ctx context.Context) ([]work.Round, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT 
		id, start_date, end_date, freeze_start_date
		FROM rounds
		ORDER BY start_date`)

	if err != nil {
		return nil, fmt.Errorf("sql error getting rounds: %w", err)
	}

	var rounds []work.Round
	for rows.Next() {
		round, err := s.scan(rows)

		if err != nil {
			return rounds, fmt.Errorf("cannot scan round: %w", err)
		}
		rounds = append(rounds, round)
	}

	return rounds, nil
}

func (s *RoundService) GetRound(ctx context.Context, roundId int) (work.Round, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT 
		id, start_date, end_date, freeze_start_date
		FROM rounds
		WHERE id = $1
		ORDER BY start_date`, roundId)

	return s.scan(row)

}

type Scanner interface {
	Scan(...any) error
}

func (s *RoundService) scan(scanner Scanner) (work.Round, error) {
	var round work.Round
	err := scanner.Scan(&round.Id, &round.StartDate, &round.EndDate, &round.FreezeStartDate)
	return round, err
}
