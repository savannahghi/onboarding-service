package mock

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/healthcrm"
)

// HealthClientMock mocks the health CRM client library implementations
type HealthClientMock struct {
	MockCreateFacilityFn func(ctx context.Context, facility *healthcrm.Facility) (*healthcrm.FacilityOutput, error)
}

// NewHealthCRMClientMock initializes our client mocks
func NewHealthCRMClientMock() *HealthClientMock {
	return &HealthClientMock{
		MockCreateFacilityFn: func(ctx context.Context, facility *healthcrm.Facility) (*healthcrm.FacilityOutput, error) {
			return &healthcrm.FacilityOutput{
				ID:            gofakeit.UUID(),
				Created:       time.Now(),
				Name:          gofakeit.BeerName(),
				Description:   gofakeit.BeerName(),
				FacilityType:  "HOSPITAL",
				County:        gofakeit.CountryAbr(),
				Country:       gofakeit.CountryAbr(),
				Coordinates:   healthcrm.CoordinatesOutput{},
				Status:        "DRAFT",
				Address:       "12-MERU",
				Contacts:      []healthcrm.ContactsOutput{},
				Identifiers:   []healthcrm.IdentifiersOutput{},
				BusinessHours: []any{},
			}, nil
		},
	}
}

// CreateFacility mocks the implementation of creating a facility
func (sc HealthClientMock) CreateFacility(ctx context.Context, facility *healthcrm.Facility) (*healthcrm.FacilityOutput, error) {
	return sc.MockCreateFacilityFn(ctx, facility)
}
