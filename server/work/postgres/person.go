package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type PersonService struct {
	db *sql.DB
}

func NewPersonService(db *sql.DB) *PersonService {
	return &PersonService{db}
}

var _ work.PersonService = (*PersonService)(nil)

func (s *PersonService) GetPerson(ctx context.Context, personID int) (work.Person, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT person_id, display_name, team FROM persons WHERE person_id = $1", personID)

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
	_, err := s.db.ExecContext(ctx, "INSERT INTO persons (person_id, display_name, team) VALUES ($1, $2, $3)", p.PersonID, p.DisplayName, p.Team)

	return err
}

func (s *PersonService) UpdatePerson(ctx context.Context, p work.Person) error {
	_, err := s.db.ExecContext(ctx, "UPDATE persons SET display_name=$2, team=$3 WHERE person_id=$1", p.PersonID, p.DisplayName, p.Team)
	return err
}

func scanPersons(rows *sql.Rows) ([]work.Person, error) {

	var persons []work.Person

	for rows.Next() {
		var p work.Person

		err := rows.Scan(&p.PersonID, &p.DisplayName, &p.Team)

		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}
