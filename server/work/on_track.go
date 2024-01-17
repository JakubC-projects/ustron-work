package work

import "context"

type OnTrackService interface {
	GetOnTrackStatus(context.Context) (Status, error)
}
