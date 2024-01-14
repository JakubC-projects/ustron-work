package mock

import (
	"context"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type DataService[T any, P comparable] struct {
	Data       []T
	PrimaryKey func(T) P
}

func NewDataService[T any, P comparable](data []T, primaryKey func(T) P) DataService[T, P] {
	return DataService[T, P]{data, primaryKey}
}

func (s *DataService[T, P]) Get(ctx context.Context, id P) (T, error) {
	res, _ := s.Find(ctx, func(d T) bool { return s.PrimaryKey(d) == id })
	if len(res) != 1 {
		var r T
		return r, work.ErrNotFound
	}
	return res[0], nil
}

func (s *DataService[T, P]) Find(ctx context.Context, predicate func(T) bool) ([]T, error) {
	var res []T
	for _, d := range s.Data {
		if predicate(d) {
			res = append(res, d)
		}
	}
	return res, nil
}

func (s *DataService[T, P]) Create(ctx context.Context, data T) error {
	s.Data = append(s.Data, data)
	return nil
}

func (s *DataService[T, P]) Update(ctx context.Context, data T) error {
	for i, d := range s.Data {
		if s.PrimaryKey(d) == s.PrimaryKey(data) {
			s.Data[i] = data
			return nil
		}
	}
	return work.ErrNotFound
}
