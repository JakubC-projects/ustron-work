package postgres

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	personService := getPersonService()

	uid := uuid.MustParse("7dbbdbd4-7347-495c-9ce9-30c5b6e51928")
	p, err := personService.GetPerson(context.Background(), uid)

	assert.NoError(t, err)
	assert.Equal(t, uid, p.Uid)
}

func TestCreateUser(t *testing.T) {
	personService := getPersonService()

	err := personService.CreatePerson(context.Background(), work.Person{
		Uid:         uuid.New(),
		DisplayName: "Test user",
		Team:        work.TeamBlue,
		Role:        work.RoleAdmin,
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
