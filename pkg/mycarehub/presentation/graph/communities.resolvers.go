package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	stream_chat "github.com/GetStream/stream-chat-go/v5"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

// CreateCommunity is the resolver for the createCommunity field.
func (r *mutationResolver) CreateCommunity(ctx context.Context, input dto.CommunityInput) (*domain.Community, error) {
	return r.mycarehub.Community.CreateCommunity(ctx, input)
}

// DeleteCommunities is the resolver for the deleteCommunities field.
func (r *mutationResolver) DeleteCommunities(ctx context.Context, communityIDs []string, hardDelete bool) (bool, error) {
	return r.mycarehub.Community.DeleteCommunities(ctx, communityIDs, hardDelete)
}

// RejectInvitation is the resolver for the rejectInvitation field.
func (r *mutationResolver) RejectInvitation(ctx context.Context, memberID string, communityID string) (bool, error) {
	return r.mycarehub.Community.RejectInvite(ctx, memberID, communityID)
}

// AcceptInvitation is the resolver for the acceptInvitation field.
func (r *mutationResolver) AcceptInvitation(ctx context.Context, memberID string, communityID string) (bool, error) {
	return r.mycarehub.Community.AcceptInvite(ctx, memberID, communityID)
}

// AddMembersToCommunity is the resolver for the addMembersToCommunity field.
func (r *mutationResolver) AddMembersToCommunity(ctx context.Context, memberIDs []string, communityID string) (bool, error) {
	return r.mycarehub.Community.AddMembersToCommunity(ctx, memberIDs, communityID)
}

// RemoveMembersFromCommunity is the resolver for the removeMembersFromCommunity field.
func (r *mutationResolver) RemoveMembersFromCommunity(ctx context.Context, communityID string, memberIDs []string) (bool, error) {
	return r.mycarehub.Community.RemoveMembersFromCommunity(ctx, communityID, memberIDs)
}

// AddModerators is the resolver for the addModerators field.
func (r *mutationResolver) AddModerators(ctx context.Context, memberIDs []string, communityID string) (bool, error) {
	return r.mycarehub.Community.AddModeratorsWithMessage(ctx, memberIDs, communityID)
}

// DemoteModerators is the resolver for the demoteModerators field.
func (r *mutationResolver) DemoteModerators(ctx context.Context, communityID string, memberIDs []string) (bool, error) {
	return r.mycarehub.Community.DemoteModerators(ctx, communityID, memberIDs)
}

// BanUser is the resolver for the banUser field.
func (r *mutationResolver) BanUser(ctx context.Context, memberID string, bannedBy string, communityID string) (bool, error) {
	return r.mycarehub.Community.BanUser(ctx, memberID, bannedBy, communityID)
}

// UnBanUser is the resolver for the unBanUser field.
func (r *mutationResolver) UnBanUser(ctx context.Context, memberID string, communityID string) (bool, error) {
	return r.mycarehub.Community.UnBanUser(ctx, memberID, communityID)
}

// DeleteCommunityMessage is the resolver for the deleteCommunityMessage field.
func (r *mutationResolver) DeleteCommunityMessage(ctx context.Context, messageID string) (bool, error) {
	return r.mycarehub.Community.DeleteCommunityMessage(ctx, messageID)
}

// ListMembers is the resolver for the listMembers field.
func (r *queryResolver) ListMembers(ctx context.Context, input *stream_chat.QueryOption) ([]*domain.Member, error) {
	return r.mycarehub.Community.ListMembers(ctx, input)
}

// ListCommunityBannedMembers is the resolver for the listCommunityBannedMembers field.
func (r *queryResolver) ListCommunityBannedMembers(ctx context.Context, communityID string) ([]*domain.Member, error) {
	return r.mycarehub.Community.ListCommunityBannedMembers(ctx, communityID)
}

// InviteMembersToCommunity is the resolver for the inviteMembersToCommunity field.
func (r *queryResolver) InviteMembersToCommunity(ctx context.Context, communityID string, memberIDs []string) (bool, error) {
	return r.mycarehub.Community.InviteMembers(ctx, communityID, memberIDs)
}

// ListCommunities is the resolver for the listCommunities field.
func (r *queryResolver) ListCommunities(ctx context.Context, input *stream_chat.QueryOption) ([]*domain.Community, error) {
	return r.mycarehub.Community.ListCommunities(ctx, input)
}

// ListCommunityMembers is the resolver for the listCommunityMembers field.
func (r *queryResolver) ListCommunityMembers(ctx context.Context, communityID string, input *stream_chat.QueryOption) ([]*domain.CommunityMember, error) {
	return r.mycarehub.Community.ListCommunityMembers(ctx, communityID, input)
}

// ListPendingInvites is the resolver for the listPendingInvites field.
func (r *queryResolver) ListPendingInvites(ctx context.Context, memberID string, input *stream_chat.QueryOption) ([]*domain.Community, error) {
	return r.mycarehub.Community.ListPendingInvites(ctx, memberID, input)
}

// RecommendedCommunities is the resolver for the recommendedCommunities field.
func (r *queryResolver) RecommendedCommunities(ctx context.Context, clientID string, limit int) ([]*domain.Community, error) {
	return r.mycarehub.Community.RecommendedCommunities(ctx, clientID, limit)
}

// ListFlaggedMessages is the resolver for the listFlaggedMessages field.
func (r *queryResolver) ListFlaggedMessages(ctx context.Context, communityCid *string, memberIDs []*string) ([]*domain.MessageFlag, error) {
	return r.mycarehub.Community.ListFlaggedMessages(ctx, communityCid, memberIDs)
}
