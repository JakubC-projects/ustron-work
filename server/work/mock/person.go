package mock

import (
	"context"

	"github.com/jakubc-projects/ustron-work/server/work"
)

type PersonService struct {
	S DataService[work.Person, int]
}

var _ work.PersonService = (*PersonService)(nil)

func NewPersonService(data ...work.Person) *PersonService {
	return &PersonService{
		S: NewDataService(data, func(d work.Person) int { return d.PersonID }),
	}
}

func (rs *PersonService) GetPerson(ctx context.Context, id int) (work.Person, error) {
	return rs.S.Get(ctx, id)
}
func (rs *PersonService) CreatePerson(ctx context.Context, reg work.Person) error {
	return rs.S.Create(ctx, reg)
}
func (rs *PersonService) UpdatePerson(ctx context.Context, reg work.Person) error {
	return rs.S.Update(ctx, reg)
}
