package work

import "context"

type Status map[Team]int

type StatusService interface {
	GetStatus(context.Context) (Status, error)
}
