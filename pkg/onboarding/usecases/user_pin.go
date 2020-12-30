package usecases

import (
	"context"
	"fmt"

	"gitlab.slade360emr.com/go/base"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/config/errors"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/config/utils"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/domain"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/repository"
)

// UserPINUseCases represents all the business logic that touch on user PIN Management
type UserPINUseCases interface {
	SetUserPIN(ctx context.Context, phone string, pin string) (*domain.PIN, error)
	ChangeUserPIN(ctx context.Context, phone string, pin string, otp string) (bool, error)
	ResetPIN(ctx context.Context, phone string) (string, error)
}

// UserPinUseCaseImpl represents usecase implementation object
type UserPinUseCaseImpl struct {
	onboardingRepository repository.OnboardingRepository
}

// NewUserPinUseCase returns a new UserPin usecase
func NewUserPinUseCase(r repository.OnboardingRepository) *UserPinUseCaseImpl {
	return &UserPinUseCaseImpl{r}
}

// SetUserPIN receives phone number and pin from phonenumber sign up
func (u *UserPinUseCaseImpl) SetUserPIN(ctx context.Context, msisdn, pin string) (*domain.PIN, error) {
	// ensure the phone number is valid
	phoneNumber, err := base.NormalizeMSISDN(msisdn)
	if err != nil {
		return nil, fmt.Errorf("unable to normalize the msisdn: %v", err)
	}

	profile, err := u.onboardingRepository.
		GetUserProfileByPrimaryPhoneNumber(ctx, msisdn)
	if err != nil {
		return nil, &domain.CustomError{
			Err:     err,
			Message: errors.ProfileNotFoundErrMsg,
			Code:    int(base.ProfileNotFound),
		}
	}
	err = utils.ValidatePINLength(pin)
	if err != nil {
		return nil, err
	}

	err = utils.ValidatePINDigits(pin)
	if err != nil {
		return nil, err
	}

	// check if user has existing PIN
	exists, err := u.CheckHasPIN(ctx, msisdn)
	if err != nil {
		return nil, fmt.Errorf("unable to check if the user has a PIN: %v", err)
	}

	// return error if the user already have one
	if exists {
		return nil, &domain.CustomError{
			Err:     err,
			Message: errors.UsePinExistErrMsg,
			// TODO: correct error code
			Code: int(base.UserNotFound),
		}
	}

	// EncryptPIN the PIN
	salt, encryptedPin := utils.EncryptPIN(pin, nil)
	if err != nil {
		return nil, &domain.CustomError{
			Err:     err,
			Message: errors.EncryptPINErrMsg,
			// TODO: correct error code
			Code: int(base.UserNotFound),
		}
	}

	pinPayload := &domain.PIN{
		ProfileID:   profile.ID,
		PhoneNumber: phoneNumber,
		PINNumber:   encryptedPin,
		Salt:        salt,
	}
	return u.onboardingRepository.SavePIN(ctx, pinPayload)
}

// CheckHasPIN given a phone number checks if the phonenumber is present in our collections
// which essentially means that the number has an already existing PIN
func (u *UserPinUseCaseImpl) CheckHasPIN(ctx context.Context, msisdn string) (bool, error) {
	uid, err := base.GetLoggedInUserUID(ctx)
	if err != nil {
		return false, fmt.Errorf("unable to get the logged in user: %v", err)
	}

	profile, err := u.onboardingRepository.GetUserProfileByUID(ctx, uid)
	if err != nil {
		return false, fmt.Errorf("unable to normalize the msisdn: %v", err)
	}

	PINData, err := u.onboardingRepository.
		GetPINByProfileID(ctx, profile.ID)

	if err != nil {
		return false, err
	}

	if PINData == nil {
		return false, &domain.CustomError{
			Err:     err,
			Message: errors.PINNotFoundErrMsg,
			Code:    int(base.PINNotFound),
		}
	}

	return true, nil
}

// ChangeUserPIN ...
func (u *UserPinUseCaseImpl) ChangeUserPIN(ctx context.Context, phone string, pin string, otp string) (bool, error) {
	return false, nil
}

// ResetPIN ...
func (u *UserPinUseCaseImpl) ResetPIN(ctx context.Context, phone string) (string, error) {
	return "", nil
}
