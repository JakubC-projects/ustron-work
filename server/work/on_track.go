package work

import "context"

type OnTrackStatus map[string]int

const (
	OnTrackMale   = "Male"
	OnTrackFemale = "Female"
	OnTrackPoland = "Poland"
)

var OnTrackEntities = append([]string{OnTrackMale, OnTrackFemale, OnTrackPoland}, Teams...)

func NewOnTrackStatus() OnTrackStatus {
	s := OnTrackStatus{}

	for _, g := range OnTrackEntities {
		s[g] = 0
	}

	return s
}

type OnTrackService interface {
	GetOnTrackStatus(ctx context.Context, roundId int) (OnTrackStatus, error)
}
