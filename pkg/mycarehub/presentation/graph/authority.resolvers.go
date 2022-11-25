package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

// AssignOrRevokeRoles is the resolver for the assignOrRevokeRoles field.
func (r *mutationResolver) AssignOrRevokeRoles(ctx context.Context, userID string, roles []*enums.UserRoleType) (bool, error) {
	return r.mycarehub.Authority.AssignOrRevokeRoles(ctx, userID, roles)
}

// GetUserRoles is the resolver for the getUserRoles field.
func (r *queryResolver) GetUserRoles(ctx context.Context, userID string, organisationID string) ([]*domain.AuthorityRole, error) {
	return r.mycarehub.Authority.GetUserRoles(ctx, userID, organisationID)
}

// GetAllAuthorityRoles is the resolver for the getAllAuthorityRoles field.
func (r *queryResolver) GetAllAuthorityRoles(ctx context.Context) ([]*domain.AuthorityRole, error) {
	return r.mycarehub.Authority.GetAllRoles(ctx)
}
