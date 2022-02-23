package communities

import (
	"context"
	"fmt"

	stream "github.com/GetStream/stream-chat-go/v5"
	"github.com/google/uuid"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/common/helpers"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/exceptions"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/extension"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure"
	streamService "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/getstream"
)

const (
	inviteMessage = "%v has invited you to join this community"
)

// ICreateCommunity is an interface that is used to create communities
type ICreateCommunity interface {
	CreateCommunity(ctx context.Context, input dto.CommunityInput) (*domain.Community, error)
}

// IListUsers is an interface that is used to list getstream users
type IListUsers interface {
	ListCommunityMembers(ctx context.Context, communityID string) ([]*domain.CommunityMember, error)
	ListMembers(ctx context.Context, input *stream.QueryOption) ([]*domain.Member, error)
}

// IInviteMembers interface holds methods that are used to send member invites
type IInviteMembers interface {
	InviteMembers(ctx context.Context, communityID string, userIDS []string) (bool, error)
}

// IListCommunities is an interface that is used to list getstream channels
type IListCommunities interface {
	ListCommunities(ctx context.Context, input *stream.QueryOption) ([]*domain.Community, error)
}

// UseCasesCommunities holds all interfaces required to implement the communities feature
type UseCasesCommunities interface {
	ICreateCommunity
	IInviteMembers
	IListUsers
	IListCommunities
}

// UseCasesCommunitiesImpl represents communities implementation
type UseCasesCommunitiesImpl struct {
	GetstreamService streamService.ServiceGetStream
	Create           infrastructure.Create
	ExternalExt      extension.ExternalMethodsExtension
	Query            infrastructure.Query
}

// NewUseCaseCommunitiesImpl initializes a new communities service
func NewUseCaseCommunitiesImpl(
	getstream streamService.ServiceGetStream,
	ext extension.ExternalMethodsExtension,
	create infrastructure.Create,
	query infrastructure.Query,
) *UseCasesCommunitiesImpl {
	return &UseCasesCommunitiesImpl{
		GetstreamService: getstream,
		Create:           create,
		ExternalExt:      ext,
		Query:            query,
	}
}

// ListMembers returns list of the members that match QueryOption that's passed as the input
func (us *UseCasesCommunitiesImpl) ListMembers(ctx context.Context, input *stream.QueryOption) ([]*domain.Member, error) {
	var query *stream.QueryOption

	if input == nil {
		query = &stream.QueryOption{
			Filter: map[string]interface{}{
				"role": "user",
			},
		}
	} else {
		query = &stream.QueryOption{
			Filter:       input.Filter,
			UserID:       input.UserID,
			Limit:        input.Limit,
			Offset:       input.Offset,
			MessageLimit: input.MessageLimit,
			MemberLimit:  input.MemberLimit,
		}
	}

	userResponse := []*domain.Member{}

	getStreamUserResponse, err := us.GetstreamService.ListGetStreamUsers(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to return getstream users :%v", err)
	}

	for _, user := range getStreamUserResponse.Users {
		var userID string
		if val, ok := user.ExtraData["userID"]; ok {
			userID = val.(string)
		}

		Users := domain.Member{
			ID:     user.ID,
			Name:   user.Name,
			Role:   user.Role,
			UserID: userID,
		}
		userResponse = append(userResponse, &Users)
	}

	return userResponse, nil
}

// CreateCommunity creates channel with the GetStream chat service
func (us *UseCasesCommunitiesImpl) CreateCommunity(ctx context.Context, input dto.CommunityInput) (*domain.Community, error) {
	channelResponse, err := us.Create.CreateChannel(ctx, &input)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, exceptions.GetLoggedInUserUIDErr(err)
	}

	loggedInUserID, err := us.ExternalExt.GetLoggedInUserUID(ctx)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, exceptions.GetLoggedInUserUIDErr(err)
	}

	data := map[string]interface{}{
		"minimumAge": channelResponse.AgeRange.LowerBound,
		"maximumAge": channelResponse.AgeRange.UpperBound,
		"gender":     channelResponse.Gender,
		"clientType": channelResponse.ClientType,
		"inviteOnly": channelResponse.InviteOnly,
		"name":       channelResponse.Name,
	}

	_, err = us.GetstreamService.CreateChannel(ctx, "messaging", channelResponse.ID, loggedInUserID, data)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, fmt.Errorf("unable to create channel: %v", err)
	}

	return &domain.Community{
		ID:          channelResponse.ID,
		CID:         channelResponse.CID,
		Name:        channelResponse.Name,
		Description: channelResponse.Description,
		AgeRange: &domain.AgeRange{
			LowerBound: channelResponse.AgeRange.LowerBound,
			UpperBound: channelResponse.AgeRange.UpperBound,
		},
		Gender:     channelResponse.Gender,
		ClientType: channelResponse.ClientType,
		InviteOnly: channelResponse.InviteOnly,
	}, nil
}

// InviteMembers invites specified members to a community
func (us *UseCasesCommunitiesImpl) InviteMembers(ctx context.Context, communityID string, userIDS []string) (bool, error) {
	loggedInUserID, err := us.ExternalExt.GetLoggedInUserUID(ctx)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.GetLoggedInUserUIDErr(err)
	}

	staffProfile, err := us.Query.GetStaffProfileByUserID(ctx, loggedInUserID)
	if err != nil {
		return false, fmt.Errorf("failed to get staff profile")
	}

	// TODO: Fetch the channel to get the channel name and pass it as part of the message
	message := &stream.Message{
		ID:   uuid.New().String(),
		Text: fmt.Sprintf(inviteMessage, staffProfile.User.Name),
		User: &stream.User{
			ID: *staffProfile.ID,
		},
	}

	_, err = us.GetstreamService.InviteMembers(ctx, userIDS, communityID, message)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to invite members to a community: %v", err)
	}

	return true, nil
}

// ListCommunities returns list of the communities that match QueryOption that's passed as the input
func (us *UseCasesCommunitiesImpl) ListCommunities(ctx context.Context, input *stream.QueryOption) ([]*domain.Community, error) {
	channelResponse := []*domain.Community{}

	var query stream.QueryOption
	// set default offset and limit if none is passed
	if input == nil {
		query = stream.QueryOption{
			Filter: map[string]interface{}{},
			Limit:  10,
			Offset: 0,
		}
	} else {
		query = stream.QueryOption{
			Filter: input.Filter,
			Limit:  input.Limit,
			Offset: input.Offset,
		}
	}

	getStreamChannelResponse, err := us.GetstreamService.ListGetStreamChannels(ctx, &query)
	if err != nil {
		return nil, fmt.Errorf("failed to return getstream channels :%v", err)
	}

	for _, channel := range getStreamChannelResponse.Channels {

		createdBy := &domain.Member{
			ID:   channel.CreatedBy.ID,
			Name: channel.CreatedBy.Name,
			Role: channel.CreatedBy.Role,
		}

		channelResponse = append(channelResponse, &domain.Community{
			ID:          channel.ID,
			CID:         channel.CID,
			CreatedBy:   createdBy,
			Disabled:    channel.Disabled,
			Frozen:      channel.Frozen,
			MemberCount: channel.MemberCount,
			CreatedAt:   channel.CreatedAt,
			UpdatedAt:   channel.UpdatedAt,
		})
	}

	return channelResponse, nil
}

// ListCommunityMembers retrieves the members of a community
func (us *UseCasesCommunitiesImpl) ListCommunityMembers(ctx context.Context, communityID string) ([]*domain.CommunityMember, error) {
	members := []*domain.CommunityMember{}

	channel, err := us.GetstreamService.GetChannel(ctx, communityID)
	if err != nil {
		return nil, err
	}

	for _, member := range channel.Members {
		var userType string
		var userID string

		if val, ok := member.User.ExtraData["userType"]; ok {
			userType = val.(string)
		}

		if val, ok := member.User.ExtraData["userID"]; ok {
			userID = val.(string)
		}

		user := domain.Member{
			ID:     member.User.ID,
			Name:   member.User.Name,
			Role:   member.User.Role,
			UserID: userID,
		}

		commMem := &domain.CommunityMember{
			UserID:      userID,
			User:        user,
			Role:        member.Role,
			IsModerator: member.IsModerator,
			UserType:    userType,
		}

		members = append(members, commMem)

	}

	return members, nil
}
