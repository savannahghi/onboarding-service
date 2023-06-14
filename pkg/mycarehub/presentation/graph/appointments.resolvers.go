package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/savannahghi/firebasetools"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/presentation/graph/generated"
	"github.com/savannahghi/scalarutils"
)

// RescheduleAppointment is the resolver for the rescheduleAppointment field.
func (r *mutationResolver) RescheduleAppointment(ctx context.Context, appointmentID string, date scalarutils.Date, caregiverID *string) (bool, error) {
	return r.mycarehub.Appointment.RescheduleClientAppointment(ctx, appointmentID, date, caregiverID)
}

// FetchClientAppointments is the resolver for the fetchClientAppointments field.
func (r *queryResolver) FetchClientAppointments(ctx context.Context, clientID string, paginationInput dto.PaginationsInput, filters []*firebasetools.FilterParam) (*domain.AppointmentsPage, error) {
	return r.mycarehub.Appointment.FetchClientAppointments(ctx, clientID, paginationInput, filters)
}

// NextRefill is the resolver for the nextRefill field.
func (r *queryResolver) NextRefill(ctx context.Context, clientID string) (*scalarutils.Date, error) {
	return r.mycarehub.Appointment.NextRefill(ctx, clientID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
