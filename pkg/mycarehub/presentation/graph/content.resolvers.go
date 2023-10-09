package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/savannahghi/feedlib"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

// ShareContent is the resolver for the shareContent field.
func (r *mutationResolver) ShareContent(ctx context.Context, input dto.ShareContentInput) (bool, error) {
	return r.mycarehub.Content.ShareContent(ctx, input)
}

// BookmarkContent is the resolver for the bookmarkContent field.
func (r *mutationResolver) BookmarkContent(ctx context.Context, clientID string, contentItemID int) (bool, error) {
	return r.mycarehub.Content.BookmarkContent(ctx, clientID, contentItemID)
}

// UnBookmarkContent is the resolver for the UnBookmarkContent field.
func (r *mutationResolver) UnBookmarkContent(ctx context.Context, clientID string, contentItemID int) (bool, error) {
	return r.mycarehub.Content.UnBookmarkContent(ctx, clientID, contentItemID)
}

// LikeContent is the resolver for the likeContent field.
func (r *mutationResolver) LikeContent(ctx context.Context, clientID string, contentID int) (bool, error) {
	r.checkPreconditions()

	return r.mycarehub.Content.LikeContent(ctx, clientID, contentID)
}

// UnlikeContent is the resolver for the unlikeContent field.
func (r *mutationResolver) UnlikeContent(ctx context.Context, clientID string, contentID int) (bool, error) {
	r.checkPreconditions()

	return r.mycarehub.Content.UnlikeContent(ctx, clientID, contentID)
}

// ViewContent is the resolver for the viewContent field.
func (r *mutationResolver) ViewContent(ctx context.Context, clientID string, contentID int) (bool, error) {
	return r.mycarehub.Content.ViewContent(ctx, clientID, contentID)
}

// GetContent is the resolver for the getContent field.
func (r *queryResolver) GetContent(ctx context.Context, categoryID *int, limit string, clientID *string) (*domain.Content, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.GetContent(ctx, categoryID, limit, clientID)
}

// ListContentCategories is the resolver for the listContentCategories field.
func (r *queryResolver) ListContentCategories(ctx context.Context) ([]*domain.ContentItemCategory, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.ListContentCategories(ctx)
}

// GetUserBookmarkedContent is the resolver for the getUserBookmarkedContent field.
func (r *queryResolver) GetUserBookmarkedContent(ctx context.Context, clientID string) (*domain.Content, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.GetUserBookmarkedContent(ctx, clientID)
}

// CheckIfUserHasLikedContent is the resolver for the checkIfUserHasLikedContent field.
func (r *queryResolver) CheckIfUserHasLikedContent(ctx context.Context, clientID string, contentID int) (bool, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.CheckWhetherUserHasLikedContent(ctx, clientID, contentID)
}

// CheckIfUserBookmarkedContent is the resolver for the checkIfUserBookmarkedContent field.
func (r *queryResolver) CheckIfUserBookmarkedContent(ctx context.Context, clientID string, contentID int) (bool, error) {
	return r.mycarehub.Content.CheckIfUserBookmarkedContent(ctx, clientID, contentID)
}

// GetFAQs is the resolver for the getFAQs field.
func (r *queryResolver) GetFAQs(ctx context.Context, flavour feedlib.Flavour) (*domain.Content, error) {
	return r.mycarehub.Content.GetFAQs(ctx, flavour)
}
