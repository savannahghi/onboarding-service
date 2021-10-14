package mock

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/savannahghi/firebasetools"
	"github.com/savannahghi/onboarding-service/pkg/onboarding/application/dto"
	"github.com/savannahghi/onboarding-service/pkg/onboarding/domain"
	"github.com/segmentio/ksuid"
	"gorm.io/datatypes"
)

// CreateMock is a mock of the create methods
type CreateMock struct {
	CreateFacilityFn func(ctx context.Context, facility dto.FacilityInput) (*domain.Facility, error)
	CollectMetricsFn func(ctx context.Context, metric *dto.MetricInput) (*domain.Metric, error)
}

// NewCreateMock initializes a new instance of `GormMock` then mocking the case of success.
func NewCreateMock() *CreateMock {
	return &CreateMock{
		CreateFacilityFn: func(ctx context.Context, facility dto.FacilityInput) (*domain.Facility, error) {
			id := uuid.New()
			name := "Kanairo One"
			code := "KN001"
			county := "Kanairo"
			description := "This is just for mocking"
			return &domain.Facility{
				ID:          id,
				Name:        name,
				Code:        code,
				Active:      true,
				County:      county,
				Description: description,
			}, nil
		},

		CollectMetricsFn: func(ctx context.Context, metric *dto.MetricInput) (*domain.Metric, error) {
			metricID := uuid.New()
			return &domain.Metric{
				MetricID:  metricID,
				Type:      domain.EngagementMetrics,
				Payload:   datatypes.JSON([]byte(`{"who": "test user", "keyword": "suicidal"}`)),
				Timestamp: time.Now(),
				UID:       ksuid.New().String(),
			}, nil
		},
	}
}

// CreateFacility mocks the implementation of `gorm's` CreateFacility method.
func (f *CreateMock) CreateFacility(ctx context.Context, facility dto.FacilityInput) (*domain.Facility, error) {
	return f.CreateFacilityFn(ctx, facility)
}

// CollectMetrics mocks the implementation of `gorm's` CollectMetrics method.
func (f *CreateMock) CollectMetrics(ctx context.Context, metric *dto.MetricInput) (*domain.Metric, error) {
	return f.CollectMetricsFn(ctx, metric)
}

// QueryMock is a mock of the query methods
type QueryMock struct {
	RetrieveFacilityFn func(ctx context.Context, id *uuid.UUID) (*domain.Facility, error)
	GetFacilitiesFn    func(ctx context.Context) ([]*domain.Facility, error)
	FindFacilityFn     func(ctx context.Context, pagination *firebasetools.PaginationInput, filter []*dto.FacilityFilterInput, sort []*dto.FacilitySortInput) (*dto.FacilityConnection, error)
}

// NewQueryMock initializes a new instance of `GormMock` then mocking the case of success.
func NewQueryMock() *QueryMock {
	return &QueryMock{

		RetrieveFacilityFn: func(ctx context.Context, id *uuid.UUID) (*domain.Facility, error) {
			facilityID := uuid.New()
			name := "test-facility"
			code := "t-100"
			county := "test-county"
			description := "test description"
			return &domain.Facility{
				ID:          facilityID,
				Name:        name,
				Code:        code,
				Active:      true,
				County:      county,
				Description: description,
			}, nil
		},
		GetFacilitiesFn: func(ctx context.Context) ([]*domain.Facility, error) {
			facilityID := uuid.New()
			name := "test-facility"
			code := "t-100"
			county := "test-county"
			description := "test description"
			return []*domain.Facility{
				{
					ID:          facilityID,
					Name:        name,
					Code:        code,
					Active:      true,
					County:      county,
					Description: description,
				},
			}, nil
		},
		FindFacilityFn: func(ctx context.Context, pagination *firebasetools.PaginationInput, filter []*dto.FacilityFilterInput, sort []*dto.FacilitySortInput) (*dto.FacilityConnection, error) {
			id := uuid.New()
			name := "Kanairo One"
			code := "KN001"
			county := "Kanairo"
			description := "This is just for mocking"

			cursor := "1"
			startCursor := "1"
			endCursor := "1"

			return &dto.FacilityConnection{
				Edges: []*dto.FacilityEdge{
					{
						Cursor: &cursor,
						Node: &domain.Facility{
							ID:          id,
							Name:        name,
							Code:        code,
							Active:      true,
							County:      county,
							Description: description,
						},
					},
				},
				PageInfo: &firebasetools.PageInfo{
					HasNextPage:     false,
					HasPreviousPage: false,
					StartCursor:     &startCursor,
					EndCursor:       &endCursor,
				},
			}, nil
		},
	}
}

// RetrieveFacility mocks the implementation of `gorm's` RetrieveFacility method.
func (f *QueryMock) RetrieveFacility(ctx context.Context, id *uuid.UUID) (*domain.Facility, error) {
	return f.RetrieveFacilityFn(ctx, id)
}

// GetFacilities mocks the implementation of `gorm's` GetFacilities method
func (f *QueryMock) GetFacilities(ctx context.Context) ([]*domain.Facility, error) {
	return f.GetFacilitiesFn(ctx)
}

// FindFacility mocks the implementation of  FindFacility method.
func (gm *QueryMock) FindFacility(ctx context.Context, pagination *firebasetools.PaginationInput, filter []*dto.FacilityFilterInput, sort []*dto.FacilitySortInput) (*dto.FacilityConnection, error) {
	return gm.FindFacilityFn(ctx, pagination, filter, sort)
}