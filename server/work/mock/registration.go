package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
)

type RegistrationService struct {
	S DataService[work.Registration, uuid.UUID]
}

var _ work.RegistrationService = (*RegistrationService)(nil)

func NewRegistrationService(data ...work.Registration) *RegistrationService {
	return &RegistrationService{
		S: NewDataService(data, func(d work.Registration) uuid.UUID { return d.Uid }),
	}
}

func (rs *RegistrationService) GetPersonRegistrations(ctx context.Context, personID int) ([]work.Registration, error) {
	return rs.S.Find(ctx, func(d work.Registration) bool { return d.PersonID == personID })
}
func (rs *RegistrationService) GetRegistration(ctx context.Context, id uuid.UUID) (work.Registration, error) {
	return rs.S.Get(ctx, id)
}
func (rs *RegistrationService) CreateRegistration(ctx context.Context, reg work.Registration) error {
	return rs.S.Create(ctx, reg)
}
func (rs *RegistrationService) UpdateRegistration(ctx context.Context, reg work.Registration) error {
	return rs.S.Update(ctx, reg)
}

func (rs *RegistrationService) GetStatus(ctx context.Context) (work.Status, error) {
	status := work.NewStatus()

	for _, r := range rs.S.Data {
		status[r.Team] += r.Value()
	}

	return status, nil

}
