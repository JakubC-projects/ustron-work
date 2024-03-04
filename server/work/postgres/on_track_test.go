package postgres

import (
	"context"
	"testing"

	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/stretchr/testify/assert"
)

func TestGetOnTrackStatus(t *testing.T) {
	onTrackService := getOnTrackService()

	reg, err := onTrackService.GetOnTrackStatus(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, reg)

	assert.Equal(t, float32(20), reg[work.TeamBlue])
}

func getOnTrackService() *OnTrackService {
	db, err := NewDb("host=localhost port=5432 user=postgres password=postgres dbname=work-test sslmode=disable")

	if err != nil {
		panic(err)
	}
	return NewOnTrackService(db)
}
