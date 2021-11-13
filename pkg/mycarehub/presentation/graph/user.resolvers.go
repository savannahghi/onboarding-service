package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

func (r *mutationResolver) AcceptTerms(ctx context.Context, userID string, termsID int) (bool, error) {
	r.checkPreconditions()
	return r.interactor.TermsUsecase.AcceptTerms(ctx, &userID, &termsID)
}

func (r *queryResolver) GetCurrentTerms(ctx context.Context) (*domain.TermsOfService, error) {
	r.checkPreconditions()
	return r.interactor.TermsUsecase.GetCurrentTerms(ctx)
}
