package work

import "context"

type OnTrackService interface {
	GetOnTrackStatus(ctx context.Context, roundId int) (Status, error)
	GetOnTrackGenderStatus(ctx context.Context, roundId int) (GenderStatus, error)
}
