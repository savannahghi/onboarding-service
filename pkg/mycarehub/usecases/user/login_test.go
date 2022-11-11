package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/feedlib"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	extensionMock "github.com/savannahghi/mycarehub/pkg/mycarehub/application/extension/mock"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	pgMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/database/postgres/mock"
	clinicalMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/clinical/mock"
	getStreamMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/getstream/mock"
	pubsubMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/pubsub/mock"
	smsMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/sms/mock"
	twilioMock "github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/services/twilio/mock"
	authorityMock "github.com/savannahghi/mycarehub/pkg/mycarehub/usecases/authority/mock"
	otpMock "github.com/savannahghi/mycarehub/pkg/mycarehub/usecases/otp/mock"
	"gorm.io/gorm"
)

func TestUseCasesUserImpl_caregiverProfileCheck(t *testing.T) {
	phone := gofakeit.Phone()
	pin := "1234"
	id := gofakeit.UUID()

	type args struct {
		ctx         context.Context
		credentials *dto.LoginInput
		response    domain.ILoginResponse
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy case: user type caregiver",
			args: args{
				ctx: context.Background(),
				credentials: &dto.LoginInput{
					PhoneNumber: &phone,
					PIN:         &pin,
					Flavour:     feedlib.FlavourConsumer,
				},
				response: &domain.LoginResponse{
					Response: &domain.Response{
						User: &domain.User{
							ID:       &id,
							Username: gofakeit.Username(),
							UserType: enums.CaregiverUser,
							Name:     gofakeit.Name(),
						},
					},
				},
			},
			want: true,
		},
		{
			name: "happy case: client that is a caregiver",
			args: args{
				ctx: context.Background(),
				credentials: &dto.LoginInput{
					PhoneNumber: &phone,
					PIN:         &pin,
					Flavour:     feedlib.FlavourConsumer,
				},
				response: &domain.LoginResponse{
					Response: &domain.Response{
						User: &domain.User{
							ID:       &id,
							Username: gofakeit.Username(),
							UserType: enums.ClientUser,
							Name:     gofakeit.Name(),
						},
					},
				},
			},
			want: true,
		},
		{
			name: "happy case: client without caregiver profile",
			args: args{
				ctx: context.Background(),
				credentials: &dto.LoginInput{
					PhoneNumber: &phone,
					PIN:         &pin,
					Flavour:     feedlib.FlavourConsumer,
				},
				response: &domain.LoginResponse{
					Response: &domain.Response{
						User: &domain.User{
							ID:       &id,
							Username: gofakeit.Username(),
							UserType: enums.ClientUser,
							Name:     gofakeit.Name(),
						},
					},
				},
			},
			want: true,
		},
		{
			name: "sad case: missing caregiver profile",
			args: args{
				ctx: context.Background(),
				credentials: &dto.LoginInput{
					PhoneNumber: &phone,
					PIN:         &pin,
					Flavour:     feedlib.FlavourConsumer,
				},
				response: &domain.LoginResponse{
					Response: &domain.Response{
						User: &domain.User{
							ID:       &id,
							Username: gofakeit.Username(),
							UserType: enums.CaregiverUser,
							Name:     gofakeit.Name(),
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeDB := pgMock.NewPostgresMock()
			fakeExtension := extensionMock.NewFakeExtension()
			fakeOTP := otpMock.NewOTPUseCaseMock()
			fakeAuthority := authorityMock.NewAuthorityUseCaseMock()
			fakeGetStream := getStreamMock.NewGetStreamServiceMock()
			fakePubsub := pubsubMock.NewPubsubServiceMock()
			fakeClinical := clinicalMock.NewClinicalServiceMock()
			fakeSMS := smsMock.NewSMSServiceMock()
			fakeTwilio := twilioMock.NewTwilioServiceMock()

			us := NewUseCasesUserImpl(fakeDB, fakeDB, fakeDB, fakeDB, fakeExtension, fakeOTP, fakeAuthority, fakeGetStream, fakePubsub, fakeClinical, fakeSMS, fakeTwilio)

			if tt.name == "happy case: client without caregiver profile" {
				fakeDB.MockGetCaregiverByUserIDFn = func(ctx context.Context, userID string) (*domain.Caregiver, error) {
					return nil, fmt.Errorf("caregiver not found: %w", gorm.ErrRecordNotFound)
				}
			}

			if tt.name == "sad case: missing caregiver profile" {
				fakeDB.MockGetCaregiverByUserIDFn = func(ctx context.Context, userID string) (*domain.Caregiver, error) {
					return nil, fmt.Errorf("caregiver not found: %w", gorm.ErrRecordNotFound)
				}
			}

			if got := us.caregiverProfileCheck(tt.args.ctx, tt.args.credentials, tt.args.response); got != tt.want {
				t.Errorf("UseCasesUserImpl.caregiverProfileCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
