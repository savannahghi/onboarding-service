package content_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	pgMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/database/postgres/mock"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/usecases/content"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/usecases/content/mock"
)

func TestUsecaseContentImpl_ListContentCategories(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*domain.ContentItemCategory
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
		{
			name: "Sad case",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			_ = mock.NewContentUsecaseMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)

			if tt.name == "Sad case" {
				fakeDB.MockListContentCategoriesFn = func(ctx context.Context) ([]*domain.ContentItemCategory, error) {
					return nil, fmt.Errorf("an error occurred")
				}
			}

			got, err := c.ListContentCategories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseContentImpl.ListContentCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && got != nil {
				t.Errorf("expected content to be nil for %v", tt.name)
				return
			}

			if !tt.wantErr && got == nil {
				t.Errorf("expected content not to be nil for %v", tt.name)
				return
			}
		})
	}
}

func TestUseCaseContentImpl_ShareContent(t *testing.T) {

	ctx := context.Background()

	type args struct {
		ctx   context.Context
		input dto.ShareContentInput
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
				ctx: ctx,
				input: dto.ShareContentInput{
					ContentID: gofakeit.Number(1, 100),
					UserID:    uuid.New().String(),
					Channel:   "SMS",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fakeDB := pgMock.NewPostgresMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)

			got, err := c.Update.ShareContent(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCaseContentImpl.ShareContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCaseContentImpl.ShareContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCasesContentImpl_BookmarkContent(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    string
		contentID int
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
				userID:    uuid.New().String(),
				contentID: gofakeit.Number(1, 44),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)
			got, err := c.BookmarkContent(tt.args.ctx, tt.args.userID, tt.args.contentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCasesContentImpl.BookmarkContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCasesContentImpl.BookmarkContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCasesContentImpl_UnBookmarkContent(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    string
		contentID int
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
				userID:    uuid.New().String(),
				contentID: gofakeit.Number(1, 44),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)
			got, err := c.UnBookmarkContent(tt.args.ctx, tt.args.userID, tt.args.contentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCasesContentImpl.UnBookmarkContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCasesContentImpl.UnBookmarkContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCasesContentImpl_GetUserBookmarkedContent(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully get user bookmarked content",
			args: args{
				ctx:    ctx,
				userID: uuid.New().String(),
			},
			wantErr: false,
		},
		{
			name: "Sad Case - Missing user ID",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
		{
			name: "Sad Case - Fail to get content",
			args: args{
				ctx:    ctx,
				userID: uuid.New().String(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			fakeContent := mock.NewContentUsecaseMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)

			if tt.name == "Sad Case - Missing user ID" {
				fakeContent.MockGetUserBookmarkedContentFn = func(ctx context.Context, userID string) (*domain.Content, error) {
					return nil, fmt.Errorf("user ID is required")
				}
			}

			if tt.name == "Sad Case - Fail to get content" {
				fakeDB.MockGetUserBookmarkedContentFn = func(ctx context.Context, userID string) ([]*domain.ContentItem, error) {
					return nil, fmt.Errorf("failed to get bookmarked content")
				}
			}

			got, err := c.GetUserBookmarkedContent(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCasesContentImpl.GetUserBookmarkedContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got %v", got)
				return
			}
		})
	}
}

func TestUseCasesContentImpl_GetContent(t *testing.T) {
	ctx := context.Background()
	categoryID := 1
	type args struct {
		ctx        context.Context
		categoryID *int
		limit      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully get content",
			args: args{
				ctx:        ctx,
				limit:      "10",
				categoryID: &categoryID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)

			got, err := c.GetContent(tt.args.ctx, tt.args.categoryID, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCasesContentImpl.GetContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got %v", got)
				return
			}
		})
	}
}

func TestUseCasesContentImpl_GetContentByContentItemID(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx       context.Context
		contentID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully get content",
			args: args{
				ctx:       ctx,
				contentID: int(uuid.New()[8]),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			c := content.NewUseCasesContentImplementation(fakeDB, fakeDB)

			got, err := c.GetContentByContentItemID(tt.args.ctx, tt.args.contentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCasesContentImpl.GetContentByContentItemID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got %v", got)
				return
			}
		})
	}
}