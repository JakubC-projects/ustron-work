package postgres

import (
	"context"
	"testing"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	personService := getPersonService()

	p, err := personService.GetPerson(context.Background(), 54512)

	assert.NoError(t, err)
	assert.Equal(t, 54512, p.PersonID)
}

func TestCreateUser(t *testing.T) {
	personService := getPersonService()

	err := personService.CreatePerson(context.Background(), work.Person{
		PersonID:    54555,
		DisplayName: "Test user",
		Team:        work.TeamBlue,
	})
	assert.NoError(t, err)
}

func getPersonService() work.PersonService {
	db, err := NewDb("host=localhost port=5432 user=postgres password=postgres dbname=test sslmode=disable")

	if err != nil {
		panic(err)
	}
	return NewPersonService(db)
}
