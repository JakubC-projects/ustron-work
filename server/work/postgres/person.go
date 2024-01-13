package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type PersonService struct {
	db *sql.DB
}

func NewPersonService(db *sql.DB) *PersonService {
	return &PersonService{db}
}

var _ work.PersonService = (*PersonService)(nil)

func (s *PersonService) GetPerson(ctx context.Context, uid uuid.UUID) (work.Person, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT uid, display_name, team, role FROM persons WHERE uid = $1", uid)

	if err != nil {
		return work.Person{}, fmt.Errorf("sql error getting person: %w", err)
	}

	persons, err := scanPersons(rows)
	if err != nil {
		return work.Person{}, fmt.Errorf("error scanning registrations: %w", err)
	}

	if len(persons) != 1 {
		return work.Person{}, work.ErrNotFound
	}

	return persons[0], nil
}

func (s *PersonService) CreatePerson(ctx context.Context, p work.Person) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO persons (uid, display_name, team, role) VALUES ($1, $2, $3, $4)", p.Uid, p.DisplayName, p.Team, p.Role)

	return err
}

func (s *PersonService) UpdatePerson(ctx context.Context, p work.Person) error {
	_, err := s.db.ExecContext(ctx, "UPDATE persons SET display_name=$2, team=$3, role=$4 WHERE uid=$1", p.Uid, p.DisplayName, p.Team, p.Role)

	return err
}

func scanPersons(rows *sql.Rows) ([]work.Person, error) {

	var persons []work.Person

	for rows.Next() {
		var p work.Person

		err := rows.Scan(&p.Uid, &p.DisplayName, &p.Team, &p.Role)

		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}
