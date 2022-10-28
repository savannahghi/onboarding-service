package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

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
func (r *mutationResolver) BookmarkContent(ctx context.Context, userID string, contentItemID int) (bool, error) {
	return r.mycarehub.Content.BookmarkContent(ctx, userID, contentItemID)
}

// UnBookmarkContent is the resolver for the UnBookmarkContent field.
func (r *mutationResolver) UnBookmarkContent(ctx context.Context, userID string, contentItemID int) (bool, error) {
	return r.mycarehub.Content.UnBookmarkContent(ctx, userID, contentItemID)
}

// LikeContent is the resolver for the likeContent field.
func (r *mutationResolver) LikeContent(ctx context.Context, userID string, contentID int) (bool, error) {
	r.checkPreconditions()

	return r.mycarehub.Content.LikeContent(ctx, userID, contentID)
}

// UnlikeContent is the resolver for the unlikeContent field.
func (r *mutationResolver) UnlikeContent(ctx context.Context, userID string, contentID int) (bool, error) {
	r.checkPreconditions()

	return r.mycarehub.Content.UnlikeContent(ctx, userID, contentID)
}

// ViewContent is the resolver for the viewContent field.
func (r *mutationResolver) ViewContent(ctx context.Context, userID string, contentID int) (bool, error) {
	return r.mycarehub.Content.ViewContent(ctx, userID, contentID)
}

// GetContent is the resolver for the getContent field.
func (r *queryResolver) GetContent(ctx context.Context, categoryID *int, limit string) (*domain.Content, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.GetContent(ctx, categoryID, limit)
}

// ListContentCategories is the resolver for the listContentCategories field.
func (r *queryResolver) ListContentCategories(ctx context.Context) ([]*domain.ContentItemCategory, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.ListContentCategories(ctx)
}

// GetUserBookmarkedContent is the resolver for the getUserBookmarkedContent field.
func (r *queryResolver) GetUserBookmarkedContent(ctx context.Context, userID string) (*domain.Content, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.GetUserBookmarkedContent(ctx, userID)
}

// CheckIfUserHasLikedContent is the resolver for the checkIfUserHasLikedContent field.
func (r *queryResolver) CheckIfUserHasLikedContent(ctx context.Context, userID string, contentID int) (bool, error) {
	r.checkPreconditions()
	return r.mycarehub.Content.CheckWhetherUserHasLikedContent(ctx, userID, contentID)
}

// CheckIfUserBookmarkedContent is the resolver for the checkIfUserBookmarkedContent field.
func (r *queryResolver) CheckIfUserBookmarkedContent(ctx context.Context, userID string, contentID int) (bool, error) {
	return r.mycarehub.Content.CheckIfUserBookmarkedContent(ctx, userID, contentID)
}

// GetFAQs is the resolver for the getFAQs field.
func (r *queryResolver) GetFAQs(ctx context.Context, flavour feedlib.Flavour) (*domain.Content, error) {
	return r.mycarehub.Content.GetFAQs(ctx, flavour)
}
