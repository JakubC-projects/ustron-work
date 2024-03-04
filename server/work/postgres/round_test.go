package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRounds(t *testing.T) {
	roundService := getRoundService()

	round, err := roundService.GetRounds(context.Background())
	assert.NoError(t, err)

	fmt.Println(round)
}

func getRoundService() *RoundService {
	db, err := NewDb("host=localhost port=5432 user=postgres password=postgres dbname=work-test sslmode=disable")

	if err != nil {
		panic(err)
	}
	return NewRoundService(db)
}
