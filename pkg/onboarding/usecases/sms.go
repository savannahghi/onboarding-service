package usecases

import (
	"context"

	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/dto"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/extension"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/utils"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/repository"
)

// SMSUsecase represent the logic involved in receiving an SMS
type SMSUsecase interface {
	CreateSMSData(ctx context.Context, input *dto.AfricasTalkingMessage) error
}

//SMSImpl represents usecase implemention object
type SMSImpl struct {
	onboardingRepository repository.OnboardingRepository
	baseExt              extension.BaseExtension
}

// NewSMSUsecase returns a new SMS usecase
func NewSMSUsecase(
	r repository.OnboardingRepository,
	ext extension.BaseExtension,
) SMSUsecase {
	return &SMSImpl{
		onboardingRepository: r,
		baseExt:              ext,
	}
}

// CreateSMSData adds SMS data of the message received
func (s *SMSImpl) CreateSMSData(ctx context.Context, input *dto.AfricasTalkingMessage) error {
	ctx, span := tracer.Start(ctx, "CreateSMSData")
	defer span.End()

	validatedInput, err := utils.ValidateAficasTalkingSMSData(input)
	if err != nil {
		utils.RecordSpanError(span, err)
		return err
	}

	err = s.onboardingRepository.PersistIncomingSMSData(ctx, validatedInput)
	if err != nil {
		utils.RecordSpanError(span, err)
		//Wrapped error, no need to wrap it again
		return err
	}

	return nil
}
