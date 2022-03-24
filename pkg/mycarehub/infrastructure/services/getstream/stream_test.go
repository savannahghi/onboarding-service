package getstream_test

import (
	"context"
	"testing"

	stream "github.com/GetStream/stream-chat-go/v5"
	"github.com/google/uuid"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/getstream"
)

func TestGetStreamClient_CreateUserGetStreamToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully generate a user token",
			args: args{
				ctx:    context.Background(),
				userID: "fe9a8f7c-f8f9-4f0c-b8b1-f8b8f8b8f8b8",
			},
			wantErr: false,
		},
		{
			name: "Sad Case - Fail to generate token",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getStream := getstream.NewServiceGetStream()
			got, err := getStream.CreateGetStreamUserToken(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStreamClient.CreateGetStreamUserToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got == "" {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}
}

func TestChatClient_ListGetStreamUsers(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *stream.QueryOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully list get stream users",
			args: args{
				ctx: context.Background(),
				input: &stream.QueryOption{
					Filter: map[string]interface{}{
						"role": "user",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad Case - Fail to get users",
			args: args{
				ctx:   context.Background(),
				input: &stream.QueryOption{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getStream := getstream.NewServiceGetStream()

			got, err := getStream.ListGetStreamUsers(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.ListGetStreamUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}
}

func TestChatClient_CreateChannel(t *testing.T) {
	g := getstream.NewServiceGetStream()

	ctx := context.Background()
	channelID := "channelJnJ"

	type args struct {
		ctx      context.Context
		chanType string
		chanID   string
		userID   string
		data     map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.CreateChannelResponse
		wantErr bool
	}{
		{
			name: "Happy case - Create channel",
			args: args{
				ctx:      ctx,
				chanType: "messaging",
				chanID:   channelID,
				userID:   "fe9a8f7c-f8f9-4f0c-b8b1-f8b8f8b8f8b8",
				data: map[string]interface{}{
					"age": map[string]interface{}{
						"lowerBound": 10,
						"upperBound": 20,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case - empty channel id",
			args: args{
				ctx:      ctx,
				chanType: "test",
				chanID:   "",
				userID:   member1,
				data:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := g.CreateChannel(tt.args.ctx, tt.args.chanType, tt.args.chanID, tt.args.userID, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}

	_, err := g.DeleteChannels(ctx, []string{"messaging:" + channelID}, true)
	if err != nil {
		t.Errorf("ChatClient.DeleteChannel() error = %v", err)
	}
}

func TestChatClient_ListGetStreamChannels(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *stream.QueryOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully list get stream channels",
			args: args{
				ctx: context.Background(),
				input: &stream.QueryOption{
					Filter: map[string]interface{}{
						"type": "messaging",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getStream := getstream.NewServiceGetStream()

			got, err := getStream.ListGetStreamChannels(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.ListGetStreamChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}
}

func TestChatClient_GetChannel(t *testing.T) {
	type args struct {
		ctx       context.Context
		channelID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case - successfully retrieve a getstream channel",
			args: args{
				ctx:       context.Background(),
				channelID: channelID,
			},
			wantErr: false,
		},
		{
			name: "Sad case - channel does not exist",
			args: args{
				ctx:       context.Background(),
				channelID: "no-existent-channel",
			},
			wantErr: true,
		},
		{
			name: "Sad case - invalid channel id",
			args: args{
				ctx:       context.Background(),
				channelID: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := c.GetChannel(tt.args.ctx, tt.args.channelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.GetChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}

func TestChatClient_RejectInvite(t *testing.T) {
	ctx := context.Background()
	user := stream.User{
		ID:        userToRejectInviteID,
		Name:      "test user accepted invite",
		Invisible: false,
	}
	streamUser, err := c.CreateGetStreamUser(ctx, &user)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	_, err = c.InviteMembers(ctx, []string{streamUser.User.ID}, ch.Channel.ID, nil)
	if err != nil {
		t.Errorf("ChatClient.InviteMembers() error = %v", err)
		return
	}
	type args struct {
		ctx       context.Context
		userID    string
		channelID string
		message   *stream.Message
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.Response
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:       ctx,
				userID:    userToRejectInviteID,
				channelID: channelID,
				message:   nil,
			},
			wantErr: false,
		},
		{
			name: "Sad case: non existent user",
			args: args{
				ctx:       ctx,
				userID:    "non-existent-user",
				channelID: channelID,
				message:   nil,
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid channel id",
			args: args{
				ctx:       ctx,
				userID:    userToRejectInviteID,
				channelID: "",
				message:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RejectInvite(tt.args.ctx, tt.args.userID, tt.args.channelID, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.RejectInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", err)
			}
		})
	}
}

func TestChatClient_AcceptInvite(t *testing.T) {
	ctx := context.Background()
	user := stream.User{
		ID:        userToAcceptInviteID,
		Name:      "test user accepted invite",
		Invisible: false,
	}
	streamUser, err := c.CreateGetStreamUser(ctx, &user)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	_, err = c.InviteMembers(ctx, []string{streamUser.User.ID}, ch.Channel.ID, nil)
	if err != nil {
		t.Errorf("ChatClient.InviteMembers() error = %v", err)
		return
	}
	customInviteMessage := "the user" + user.Name + "accepted the invite"
	type args struct {
		ctx       context.Context
		userID    string
		channelID string
		message   *stream.Message
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.Response
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:       ctx,
				userID:    userToAcceptInviteID,
				channelID: channelID,
				message: &stream.Message{
					Text: customInviteMessage,
					User: &stream.User{
						ID:   user.ID,
						Name: user.Name,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: user does not exist",
			args: args{
				ctx:       ctx,
				userID:    "no-existent-user",
				channelID: channelID,
				message:   nil,
			},
			wantErr: true,
		},
		{
			name: "Sad case: channel does not exist",
			args: args{
				ctx:       ctx,
				userID:    userToAcceptInviteID,
				channelID: "no-existent-channel",
				message:   &stream.Message{Text: customInviteMessage},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.AcceptInvite(tt.args.ctx, tt.args.userID, tt.args.channelID, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.AcceptInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("ChatClient.AcceptInvite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatClient_RemoveMembers(t *testing.T) {
	ctx := context.Background()
	_, err := c.AddMembersToCommunity(ctx, []string{member1}, ch.Channel.ID)
	if err != nil {
		t.Errorf("ChatClient.AddMembersToCommunity() error = %v", err)
		return
	}
	type args struct {
		ctx       context.Context
		channelID string
		memberIDs []string
		message   *stream.Message
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.Response
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:       ctx,
				channelID: channelID,
				memberIDs: []string{member1},
				message:   nil,
			},
			wantErr: false,
		},
		{
			name: "Sad case: channel does not exist",
			args: args{
				ctx:       ctx,
				channelID: "no-existent-channel",
				memberIDs: []string{member1},
				message:   nil,
			},
			wantErr: true,
		},
		{
			name: "Sad case: non-existent member",
			args: args{
				ctx:       ctx,
				channelID: channelID,
				memberIDs: []string{"no-existent-member"},
				message:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RemoveMembersFromCommunity(tt.args.ctx, tt.args.channelID, tt.args.memberIDs, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.RemoveMembersFromCommunity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", err)
			}
		})
	}
}

func TestChatClient_DemoteModerators(t *testing.T) {
	ctx := context.Background()
	moderatorID := "5c272749-8ad7-4b1b-b581-302d85833f4f"
	moderator1 := stream.User{
		ID:        moderatorID,
		Name:      "test user demoted moderator",
		Invisible: false,
	}

	_, err := c.CreateGetStreamUser(ctx, &moderator1)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	_, err = c.AddModeratorsWithMessage(ctx, []string{moderatorID}, channelID, nil)
	if err != nil {
		t.Errorf("ChatClient.AddModeratorsWithMessage() error = %v", err)
		return
	}

	type args struct {
		ctx       context.Context
		channelID string
		memberIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.Response
		wantErr bool
	}{
		{
			name: "Happy case: demote moderators",
			args: args{
				ctx:       ctx,
				channelID: channelID,
				memberIDs: []string{moderatorID},
			},
			wantErr: false,
		},
		{
			name: "Sad case: non-existent channel",
			args: args{
				ctx:       ctx,
				channelID: "no-existent-channel",
				memberIDs: []string{moderatorID},
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid user id",
			args: args{
				ctx:       ctx,
				channelID: channelID,
				memberIDs: []string{"no-existent-member"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.DemoteModerators(tt.args.ctx, tt.args.channelID, tt.args.memberIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.DemoteModerators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", err)
			}
		})
	}

	// teardown
	_, err = c.DeleteUsers(ctx, []string{moderatorID}, stream.DeleteUserOptions{
		User:     stream.HardDelete,
		Messages: stream.HardDelete,
	})
	if err != nil {
		t.Errorf("ChatClient.DeleteUsers() error = %v", err)
		return
	}
}

func TestChatClient_RevokeGetStreamUserToken(t *testing.T) {
	ctx := context.Background()

	streamUser := &stream.User{
		ID:        uuid.New().String(),
		Name:      "Test",
		Role:      "moderator",
		Invisible: false,
	}

	user, err := c.CreateGetStreamUser(ctx, streamUser)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	_, err = c.CreateGetStreamUserToken(ctx, user.User.ID)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUserToken() error = %v", err)
		return
	}

	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *stream.Response
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:    ctx,
				userID: user.User.ID,
			},
			wantErr: false,
		},
		{
			name: "Sad case",
			args: args{
				ctx:    ctx,
				userID: uuid.New().String(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RevokeGetStreamUserToken(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.RevokeGetStreamUserToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", err)
			}
		})
	}
	// teardown
	_, err = c.DeleteUsers(ctx, []string{user.User.ID}, stream.DeleteUserOptions{
		User:     stream.HardDelete,
		Messages: stream.HardDelete,
	})
	if err != nil {
		t.Errorf("ChatClient.DeleteUsers() error = %v", err)
		return
	}
}

func TestChatClient_BanUser(t *testing.T) {
	ctx := context.Background()

	// Create the user
	streamUser := &stream.User{
		ID:        uuid.New().String(),
		Name:      "Test",
		Invisible: false,
	}

	user, err := c.CreateGetStreamUser(ctx, streamUser)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	// Add them to a community
	_, err = c.AddMembersToCommunity(ctx, []string{user.User.ID}, channelID)
	if err != nil {
		t.Errorf("ChatClient.AddMembersToCommunity() error = %v", err)
		return
	}

	type args struct {
		ctx            context.Context
		targetMemberID string
		bannedBy       string
		communityID    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx:            ctx,
				targetMemberID: user.User.ID,
				bannedBy:       member2,
				communityID:    ch.Channel.ID,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Sad Case",
			args: args{
				ctx:            ctx,
				targetMemberID: "",
				bannedBy:       "",
				communityID:    "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Sad Case - same targetMemberID",
			args: args{
				ctx:            ctx,
				targetMemberID: user.User.ID,
				bannedBy:       user.User.ID,
				communityID:    ch.Channel.ID,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.BanUser(tt.args.ctx, tt.args.targetMemberID, tt.args.bannedBy, tt.args.communityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.BanUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChatClient.BanUser() = %v, want %v", got, tt.want)
			}
		})
	}

	// teardown
	_, err = c.DeleteUsers(ctx, []string{user.User.ID}, stream.DeleteUserOptions{
		User:     stream.HardDelete,
		Messages: stream.HardDelete,
	})
	if err != nil {
		t.Errorf("ChatClient.DeleteUsers() error = %v", err)
		return
	}
	_, err = c.DeleteChannels(ctx, []string{"messaging:" + channelID}, true)
	if err != nil {
		t.Errorf("ChatClient.DeleteChannel() error = %v", err)
	}
}

func TestChatClient_UnBanUser(t *testing.T) {
	ctx := context.Background()

	streamUser := &stream.User{
		ID:        uuid.New().String(),
		Name:      "Test",
		Invisible: false,
	}

	user, err := c.CreateGetStreamUser(ctx, streamUser)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	channel, err := c.CreateChannel(ctx, "messaging", "channelToUnbanFrom", user.User.ID, nil)
	if err != nil {
		t.Errorf("ChatClient.CreateChannel() error = %v", err)
		return
	}

	// Add them to a community
	_, err = c.AddMembersToCommunity(ctx, []string{user.User.ID, member2}, channel.Channel.ID)
	if err != nil {
		t.Errorf("ChatClient.AddMembersToCommunity() error = %v", err)
		return
	}

	_, err = c.BanUser(ctx, user.User.ID, member2, channel.Channel.ID)
	if err != nil {
		t.Errorf("unable to ban user: %v", err)
		return
	}

	type args struct {
		ctx         context.Context
		targetID    string
		communityID string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:         ctx,
				targetID:    user.User.ID,
				communityID: channel.Channel.ID,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Sad case",
			args: args{
				ctx:         ctx,
				targetID:    uuid.New().String(),
				communityID: channel.Channel.ID,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.UnBanUser(tt.args.ctx, tt.args.targetID, tt.args.communityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.UnBanUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChatClient.UnBanUser() = %v, want %v", got, tt.want)
			}
		})
	}

	// teardown
	_, err = c.DeleteUsers(ctx, []string{user.User.ID}, stream.DeleteUserOptions{
		User:     stream.HardDelete,
		Messages: stream.HardDelete,
	})
	if err != nil {
		t.Errorf("ChatClient.DeleteUsers() error = %v", err)
		return
	}
	_, err = c.DeleteChannels(ctx, []string{"messaging:" + channel.Channel.CID}, true)
	if err != nil {
		t.Errorf("ChatClient.DeleteChannel() error = %v", err)
	}
}

func TestChatClient_ListCommunityBannedMembers(t *testing.T) {
	ctx := context.Background()

	streamUser := &stream.User{
		ID:        uuid.New().String(),
		Name:      "Test",
		Invisible: false,
	}

	user, err := c.CreateGetStreamUser(ctx, streamUser)
	if err != nil {
		t.Errorf("ChatClient.CreateGetStreamUser() error = %v", err)
		return
	}

	channel, err := c.CreateChannel(ctx, "messaging", "channelToListMembersFrom", user.User.ID, nil)
	if err != nil {
		t.Errorf("ChatClient.CreateChannel() error = %v", err)
		return
	}

	// Add them to a community
	_, err = c.AddMembersToCommunity(ctx, []string{user.User.ID, member2}, channel.Channel.ID)
	if err != nil {
		t.Errorf("ChatClient.AddMembersToCommunity() error = %v", err)
		return
	}

	_, err = c.BanUser(ctx, user.User.ID, member2, channel.Channel.ID)
	if err != nil {
		t.Errorf("unable to ban user: %v", err)
		return
	}

	type args struct {
		ctx         context.Context
		communityID string
	}
	tests := []struct {
		name    string
		args    args
		want    []*domain.Member
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:         ctx,
				communityID: channel.Channel.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListCommunityBannedMembers(tt.args.ctx, tt.args.communityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.ListCommunityBannedMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	// teardown
	_, err = c.DeleteUsers(ctx, []string{user.User.ID}, stream.DeleteUserOptions{
		User:     stream.HardDelete,
		Messages: stream.HardDelete,
	})
	if err != nil {
		t.Errorf("ChatClient.DeleteUsers() error = %v", err)
		return
	}
	_, err = c.DeleteChannels(ctx, []string{"messaging:" + channel.Channel.ID}, true)
	if err != nil {
		t.Errorf("ChatClient.DeleteChannel() error = %v", err)
	}
}

func TestChatClient_UpsertUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *stream.User
	}
	tests := []struct {
		name string

		args    args
		want    *stream.UpsertUserResponse
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx: context.Background(),
				user: &stream.User{
					ID:   member1,
					Name: "Test",
					Role: "moderator",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.UpsertUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatClient.UpsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("ChatClient.UpsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
