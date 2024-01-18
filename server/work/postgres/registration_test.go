package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/jakubc-projects/ustron-work/server/work/date"
	"github.com/stretchr/testify/assert"
)

func TestCreateRegistration(t *testing.T) {
	registrationService := getRegistrationService()

	err := registrationService.CreateRegistration(context.Background(), work.Registration{
		PersonID: 54513,
		Uid:      uuid.New(),
		Team:     work.TeamBlue,
		Type:     work.RegistrationTypeMoney,
		PaidSum:  1000,
		Date:     date.DateOf(time.Now()),
		Comment:  "Hello world",
	})
	assert.NoError(t, err)
}

func TestGetRegistration(t *testing.T) {
	registrationService := getRegistrationService()

	uid := uuid.MustParse("043f8184-64d1-410a-a38d-e66b9e9fc230")
	reg, err := registrationService.GetRegistration(context.Background(), uid)

	assert.Equal(t, uid, reg.Uid)
	assert.NoError(t, err)
}

func getRegistrationService() *RegistrationService {
	db, err := NewDb("host=localhost port=5432 user=postgres password=postgres dbname=work-test sslmode=disable")

	if err != nil {
		panic(err)
	}
	return NewRegistrationService(db)
}
