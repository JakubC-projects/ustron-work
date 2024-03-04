package work

import (
	"context"
	"database/sql"
	"time"
)

type Round struct {
	Id              int          `json:"id"`
	StartDate       time.Time    `json:"startDate"`
	EndDate         time.Time    `json:"endDate"`
	FreezeStartDate sql.NullTime `json:"freezeStartDate"`
}

type RoundService interface {
	GetRounds(context.Context) ([]Round, error)
	GetRound(context.Context, int) (Round, error)
}

func GetCurrentRound(rounds []Round) (Round, bool) {
	now := time.Now()
	for _, r := range rounds {
		if r.StartDate.Before(now) && r.EndDate.After(now) {
			return r, true
		}
	}
	return Round{}, false
}
