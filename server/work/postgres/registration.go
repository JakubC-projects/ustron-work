package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type RegistrationService struct {
	db *sql.DB
}

func NewRegistrationService(db *sql.DB) *RegistrationService {
	return &RegistrationService{db}
}

var _ work.RegistrationService = (*RegistrationService)(nil)

func (s *RegistrationService) GetRegistration(ctx context.Context, uid uuid.UUID) (work.Registration, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT 
		uid, person_id, team, type, hourly_wage, hours, paid_sum
		FROM Registrations WHERE uid = $1`, uid)

	if err != nil {
		return work.Registration{}, fmt.Errorf("sql error getting Registration: %w", err)
	}

	regs, err := scanRegistrations(rows)
	if err != nil {
		return work.Registration{}, fmt.Errorf("error scanning registrations: %w", err)
	}

	if len(regs) != 1 {
		return work.Registration{}, work.ErrNotFound
	}

	return regs[0], nil
}

func (s *RegistrationService) GetPersonRegistrations(ctx context.Context, personID int) ([]work.Registration, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT 
		uid, person_id, team, type, hourly_wage, hours, paid_sum
		FROM registrations WHERE person_id = $1`, personID)

	if err != nil {
		return nil, fmt.Errorf("sql error getting registrations: %w", err)
	}

	return scanRegistrations(rows)
}

func (s *RegistrationService) GetTeamRegistrations(ctx context.Context, team work.Team) ([]work.Registration, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT 
		uid, person_id, team, type, hourly_wage, hours, paid_sum
		FROM registrations WHERE team = $1`, team)
	if err != nil {
		return nil, fmt.Errorf("sql error getting Registration: %w", err)
	}

	return scanRegistrations(rows)

}

func (s *RegistrationService) CreateRegistration(ctx context.Context, r work.Registration) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO registrations (uid, person_id, team, type, hourly_wage, hours, paid_sum) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		r.Uid, r.PersonID, r.Team, r.Type, r.HourlyWage, r.Hours, r.PaidSum)

	return err
}

func (s *RegistrationService) UpdateRegistration(ctx context.Context, r work.Registration) error {
	_, err := s.db.ExecContext(ctx, "UPDATE registrations SET person_id=$2, team=$3, type=$4 hourly_wage=$5, hours=$6, paid_sum=$7 WHERE uid=$1",
		r.Uid, r.PersonID, r.Team, r.Type, r.HourlyWage, r.Hours, r.PaidSum)

	return err
}

func (s *RegistrationService) GetStatus(ctx context.Context) (work.Status, error) {
	result := work.NewStatus()

	rows, err := s.db.QueryContext(ctx,
		`SELECT team, SUM(paid_sum + (hourly_wage * hours)) 
		FROM registrations GROUP BY team`,
	)
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

func scanRegistrations(rows *sql.Rows) ([]work.Registration, error) {

	var registrations []work.Registration

	for rows.Next() {
		var r work.Registration
		err := rows.Scan(&r.Uid, &r.PersonID, &r.Team, &r.Type, &r.HourlyWage, &r.Hours, &r.PaidSum)

		if err != nil {
			return nil, err
		}
		registrations = append(registrations, r)
	}
	return registrations, nil
}
