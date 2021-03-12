package mock

import (
	"context"

	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/resources"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/domain"

	"gitlab.slade360emr.com/go/base"
)

// FakeOnboardingRepository is a mock onboarding repository.
type FakeOnboardingRepository struct {
	GetSupplierProfileByIDFn        func(ctx context.Context, id string) (*base.Supplier, error)
	GetSupplierProfileByUIDFn       func(ctx context.Context, uid string) (*base.Supplier, error)
	GetSupplierProfileByProfileIDFn func(ctx context.Context, profileID string) (*base.Supplier, error)
	AddPartnerTypeFn                func(ctx context.Context, profileID string, name *string, partnerType *base.PartnerType) (bool, error)

	UpdateSupplierProfileFn  func(ctx context.Context, profileID string, data *base.Supplier) error
	AddSupplierAccountTypeFn func(ctx context.Context, profileID string, accountType base.AccountType) (*base.Supplier, error)

	StageProfileNudgeFn func(ctx context.Context, nudge *base.Nudge) error

	StageKYCProcessingRequestFn func(ctx context.Context, data *domain.KYCRequest) error

	RemoveKYCProcessingRequestFn func(ctx context.Context, supplierProfileID string) error

	// sets the active attribute of supplier profile to true
	ActivateSupplierProfileFn func(profileID string, supplier base.Supplier) (*base.Supplier, error)

	FetchKYCProcessingRequestsFn func(ctx context.Context) ([]*domain.KYCRequest, error)

	FetchKYCProcessingRequestByIDFn func(ctx context.Context, id string) (*domain.KYCRequest, error)

	UpdateKYCProcessingRequestFn func(ctx context.Context, sup *domain.KYCRequest) error
	GetCustomerProfileByIDFn     func(ctx context.Context, id string) (*base.Customer, error)

	GetCustomerProfileByProfileIDFn func(ctx context.Context, profileID string) (*base.Customer, error)

	CreateUserProfileFn func(ctx context.Context, phoneNumber, uid string) (*base.UserProfile, error)

	// creates an empty supplier profile
	CreateEmptySupplierProfileFn func(ctx context.Context, profileID string) (*base.Supplier, error)

	// creates an empty customer profile
	CreateEmptyCustomerProfileFn func(ctx context.Context, profileID string) (*base.Customer, error)

	// fetches a user profile by uid
	GetUserProfileByUIDFn func(ctx context.Context, uid string, suspended bool) (*base.UserProfile, error)

	// fetches a user profile by id
	GetUserProfileByIDFn func(ctx context.Context, id string, suspended bool) (*base.UserProfile, error)

	// fetches a user profile by phone number
	GetUserProfileByPhoneNumberFn func(ctx context.Context, phoneNumber string, suspended bool) (*base.UserProfile, error)

	// fetches a user profile by primary phone number
	GetUserProfileByPrimaryPhoneNumberFn func(ctx context.Context, phoneNumber string, suspended bool) (*base.UserProfile, error)

	// checks if a specific phone number has already been registered to another user
	CheckIfPhoneNumberExistsFn func(ctx context.Context, phone string) (bool, error)

	CheckIfEmailExistsFn func(ctx context.Context, email string) (bool, error)

	// checks if a specific username has already been registered to another user
	CheckIfUsernameExistsFn func(ctx context.Context, phone string) (bool, error)

	GenerateAuthCredentialsForAnonymousUserFn func(ctx context.Context) (*base.AuthCredentialResponse, error)

	GenerateAuthCredentialsFn func(ctx context.Context, phone string) (*base.AuthCredentialResponse, error)

	FetchAdminUsersFn func(ctx context.Context) ([]*base.UserProfile, error)

	// removes user completely. This should be used only under testing environment
	PurgeUserByPhoneNumberFn func(ctx context.Context, phone string) error

	HardResetSecondaryPhoneNumbersFn func(ctx context.Context, id string, phoneNumbers []string) error

	HardResetSecondaryEmailAddressFn func(ctx context.Context, id string, newSecondaryEmails []string) error

	// PINs
	GetPINByProfileIDFn func(ctx context.Context, ProfileID string) (*domain.PIN, error)

	// Record post visit survey
	RecordPostVisitSurveyFn func(ctx context.Context, input resources.PostVisitSurveyInput, UID string) error

	// User Pin methods
	SavePINFn   func(ctx context.Context, pin *domain.PIN) (bool, error)
	UpdatePINFn func(ctx context.Context, id string, pin *domain.PIN) (bool, error)

	ExchangeRefreshTokenForIDTokenFn func(
		token string,
	) (*base.AuthCredentialResponse, error)

	GetCustomerOrSupplierProfileByProfileIDFn func(
		ctx context.Context,
		flavour base.Flavour,
		profileID string,
	) (*base.Customer, *base.Supplier, error)

	GetOrCreatePhoneNumberUserFn func(
		ctx context.Context,
		phone string,
	) (*resources.CreatedUserResponse, error)

	GetUserProfileAttributesFn func(
		ctx context.Context,
		UIDs []string,
		attribute string,
	) (map[string][]string, error)

	CheckIfExperimentParticipantFn func(ctx context.Context, profileID string) (bool, error)

	AddUserAsExperimentParticipantFn func(ctx context.Context, profile *base.UserProfile) (bool, error)

	RemoveUserAsExperimentParticipantFn func(ctx context.Context, profile *base.UserProfile) (bool, error)

	AddNHIFDetailsFn func(
		ctx context.Context,
		input resources.NHIFDetailsInput,
		profileID string,
	) (*domain.NHIFDetails, error)

	GetNHIFDetailsByProfileIDFn func(
		ctx context.Context,
		profileID string,
	) (*domain.NHIFDetails, error)

	GetUserCommunicationsSettingsFn func(ctx context.Context, profileID string) (*base.UserCommunicationsSetting, error)

	SetUserCommunicationsSettingsFn func(ctx context.Context, profileID string,
		allowWhatsApp *bool, allowTextSms *bool, allowPush *bool, allowEmail *bool) (*base.UserCommunicationsSetting, error)

	UpdateCustomerProfileFn func(
		ctx context.Context,
		profileID string,
		customer base.Customer,
	) (*base.Customer, error)

	// Userprofile
	UpdateUserNameFn                func(ctx context.Context, id string, phoneNumber string) error
	UpdatePrimaryPhoneNumberFn      func(ctx context.Context, id string, phoneNumber string) error
	UpdatePrimaryEmailAddressFn     func(ctx context.Context, id string, emailAddress string) error
	UpdateSecondaryPhoneNumbersFn   func(ctx context.Context, id string, phoneNumbers []string) error
	UpdateSecondaryEmailAddressesFn func(ctx context.Context, id string, emailAddresses []string) error
	UpdateSuspendedFn               func(ctx context.Context, id string, status bool) error
	UpdatePhotoUploadIDFn           func(ctx context.Context, id string, uploadID string) error
	UpdateCoversFn                  func(ctx context.Context, id string, covers []base.Cover) error
	UpdatePushTokensFn              func(ctx context.Context, id string, pushToken []string) error
	UpdatePermissionsFn             func(ctx context.Context, id string, perms []base.PermissionType) error
	UpdateBioDataFn                 func(ctx context.Context, id string, data base.BioData) error
	UpdateVerifiedIdentifiersFn     func(ctx context.Context, id string, identifiers []base.VerifiedIdentifier) error
	UpdateVerifiedUIDSFn            func(ctx context.Context, id string, uids []string) error
	UpdateAddressesFn               func(ctx context.Context, id string, address base.Address, addressType base.AddressType) error
}

// GetSupplierProfileByID ...
func (f *FakeOnboardingRepository) GetSupplierProfileByID(ctx context.Context, id string) (*base.Supplier, error) {
	return f.GetSupplierProfileByIDFn(ctx, id)
}

// GetSupplierProfileByUID ...
func (f *FakeOnboardingRepository) GetSupplierProfileByUID(ctx context.Context, uid string) (*base.Supplier, error) {
	return f.GetSupplierProfileByUIDFn(ctx, uid)
}

// GetSupplierProfileByProfileID ...
func (f *FakeOnboardingRepository) GetSupplierProfileByProfileID(ctx context.Context, profileID string) (*base.Supplier, error) {
	return f.GetSupplierProfileByProfileIDFn(ctx, profileID)
}

// AddPartnerType ...
func (f *FakeOnboardingRepository) AddPartnerType(ctx context.Context, profileID string, name *string, partnerType *base.PartnerType) (bool, error) {
	return f.AddPartnerTypeFn(ctx, profileID, name, partnerType)
}

// UpdateSupplierProfile ...
func (f *FakeOnboardingRepository) UpdateSupplierProfile(ctx context.Context, profileID string, data *base.Supplier) error {
	return f.UpdateSupplierProfileFn(ctx, profileID, data)
}

// AddSupplierAccountType ...
func (f *FakeOnboardingRepository) AddSupplierAccountType(ctx context.Context, profileID string, accountType base.AccountType) (*base.Supplier, error) {
	return f.AddSupplierAccountTypeFn(ctx, profileID, accountType)
}

// StageProfileNudge ...
func (f *FakeOnboardingRepository) StageProfileNudge(ctx context.Context, nudge *base.Nudge) error {
	return f.StageProfileNudgeFn(ctx, nudge)
}

// StageKYCProcessingRequest ...
func (f *FakeOnboardingRepository) StageKYCProcessingRequest(ctx context.Context, data *domain.KYCRequest) error {
	return f.StageKYCProcessingRequestFn(ctx, data)
}

// RemoveKYCProcessingRequest ...
func (f *FakeOnboardingRepository) RemoveKYCProcessingRequest(ctx context.Context, supplierProfileID string) error {
	return f.RemoveKYCProcessingRequestFn(ctx, supplierProfileID)
}

// ActivateSupplierProfile ...
func (f *FakeOnboardingRepository) ActivateSupplierProfile(profileID string, supplier base.Supplier) (*base.Supplier, error) {
	return f.ActivateSupplierProfileFn(profileID, supplier)
}

// FetchKYCProcessingRequests ...
func (f *FakeOnboardingRepository) FetchKYCProcessingRequests(ctx context.Context) ([]*domain.KYCRequest, error) {
	return f.FetchKYCProcessingRequestsFn(ctx)
}

// FetchKYCProcessingRequestByID ...
func (f *FakeOnboardingRepository) FetchKYCProcessingRequestByID(ctx context.Context, id string) (*domain.KYCRequest, error) {
	return f.FetchKYCProcessingRequestByIDFn(ctx, id)
}

// UpdateKYCProcessingRequest ...
func (f *FakeOnboardingRepository) UpdateKYCProcessingRequest(ctx context.Context, sup *domain.KYCRequest) error {
	return f.UpdateKYCProcessingRequestFn(ctx, sup)
}

// GetCustomerProfileByID ...
func (f *FakeOnboardingRepository) GetCustomerProfileByID(ctx context.Context, id string) (*base.Customer, error) {
	return f.GetCustomerProfileByIDFn(ctx, id)
}

// GetCustomerProfileByProfileID ...
func (f *FakeOnboardingRepository) GetCustomerProfileByProfileID(ctx context.Context, profileID string) (*base.Customer, error) {
	return f.GetCustomerProfileByProfileIDFn(ctx, profileID)
}

// CreateUserProfile ...
func (f *FakeOnboardingRepository) CreateUserProfile(ctx context.Context, phoneNumber, uid string) (*base.UserProfile, error) {
	return f.CreateUserProfileFn(ctx, phoneNumber, uid)
}

// CreateEmptySupplierProfile ...
func (f *FakeOnboardingRepository) CreateEmptySupplierProfile(ctx context.Context, profileID string) (*base.Supplier, error) {
	return f.CreateEmptySupplierProfileFn(ctx, profileID)
}

// CreateEmptyCustomerProfile creates an empty customer profile
func (f *FakeOnboardingRepository) CreateEmptyCustomerProfile(ctx context.Context, profileID string) (*base.Customer, error) {
	return f.CreateEmptyCustomerProfileFn(ctx, profileID)
}

// GetUserProfileByUID fetches a user profile by uidActivateSupplierProfile
func (f *FakeOnboardingRepository) GetUserProfileByUID(ctx context.Context, uid string, suspended bool) (*base.UserProfile, error) {
	return f.GetUserProfileByUIDFn(ctx, uid, suspended)
}

// GetUserProfileByID fetches a user profile by id
func (f *FakeOnboardingRepository) GetUserProfileByID(ctx context.Context, id string, suspended bool) (*base.UserProfile, error) {
	return f.GetUserProfileByIDFn(ctx, id, suspended)
}

// GetUserProfileByPhoneNumber fetches a user profile by phone number
func (f *FakeOnboardingRepository) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, suspended bool) (*base.UserProfile, error) {
	return f.GetUserProfileByPhoneNumberFn(ctx, phoneNumber, suspended)
}

// GetUserProfileByPrimaryPhoneNumber fetches a user profile by primary phone number
func (f *FakeOnboardingRepository) GetUserProfileByPrimaryPhoneNumber(ctx context.Context, phoneNumber string, suspended bool) (*base.UserProfile, error) {
	return f.GetUserProfileByPrimaryPhoneNumberFn(ctx, phoneNumber, suspended)
}

// CheckIfPhoneNumberExists checks if a specific phone number has already been registered to another user
func (f *FakeOnboardingRepository) CheckIfPhoneNumberExists(ctx context.Context, phone string) (bool, error) {
	return f.CheckIfPhoneNumberExistsFn(ctx, phone)
}

// CheckIfEmailExists ...
func (f *FakeOnboardingRepository) CheckIfEmailExists(ctx context.Context, email string) (bool, error) {
	return f.CheckIfEmailExistsFn(ctx, email)
}

// CheckIfUsernameExists checks if a specific username has already been registered to another user
func (f *FakeOnboardingRepository) CheckIfUsernameExists(ctx context.Context, phone string) (bool, error) {
	return f.CheckIfUsernameExistsFn(ctx, phone)
}

// GenerateAuthCredentialsForAnonymousUser ...
func (f *FakeOnboardingRepository) GenerateAuthCredentialsForAnonymousUser(ctx context.Context) (*base.AuthCredentialResponse, error) {
	return f.GenerateAuthCredentialsForAnonymousUserFn(ctx)
}

// GenerateAuthCredentials ...
func (f *FakeOnboardingRepository) GenerateAuthCredentials(ctx context.Context, phone string) (*base.AuthCredentialResponse, error) {
	return f.GenerateAuthCredentialsFn(ctx, phone)
}

// FetchAdminUsers ...
func (f *FakeOnboardingRepository) FetchAdminUsers(ctx context.Context) ([]*base.UserProfile, error) {
	return f.FetchAdminUsersFn(ctx)
}

// PurgeUserByPhoneNumber removes user completely. This should be used only under testing environment
func (f *FakeOnboardingRepository) PurgeUserByPhoneNumber(ctx context.Context, phone string) error {
	return f.PurgeUserByPhoneNumberFn(ctx, phone)
}

// GetPINByProfileID PINs
func (f *FakeOnboardingRepository) GetPINByProfileID(ctx context.Context, ProfileID string) (*domain.PIN, error) {
	return f.GetPINByProfileIDFn(ctx, ProfileID)
}

//RecordPostVisitSurvey Record post visit survey
func (f *FakeOnboardingRepository) RecordPostVisitSurvey(ctx context.Context, input resources.PostVisitSurveyInput, UID string) error {
	return f.RecordPostVisitSurveyFn(ctx, input, UID)
}

//SavePIN  User Pin methods
func (f *FakeOnboardingRepository) SavePIN(ctx context.Context, pin *domain.PIN) (bool, error) {
	return f.SavePINFn(ctx, pin)
}

// UpdatePIN ...
func (f *FakeOnboardingRepository) UpdatePIN(ctx context.Context, id string, pin *domain.PIN) (bool, error) {
	return f.UpdatePINFn(ctx, id, pin)
}

// ExchangeRefreshTokenForIDToken ...
func (f *FakeOnboardingRepository) ExchangeRefreshTokenForIDToken(token string) (*base.AuthCredentialResponse, error) {
	return f.ExchangeRefreshTokenForIDTokenFn(token)
}

// GetCustomerOrSupplierProfileByProfileID ...
func (f *FakeOnboardingRepository) GetCustomerOrSupplierProfileByProfileID(ctx context.Context, flavour base.Flavour, profileID string) (*base.Customer, *base.Supplier, error) {
	return f.GetCustomerOrSupplierProfileByProfileIDFn(ctx, flavour, profileID)
}

// UpdateUserName ...
func (f *FakeOnboardingRepository) UpdateUserName(ctx context.Context, id string, phoneNumber string) error {
	return f.UpdateUserNameFn(ctx, id, phoneNumber)
}

// UpdatePrimaryPhoneNumber ...
func (f *FakeOnboardingRepository) UpdatePrimaryPhoneNumber(ctx context.Context, id string, phoneNumber string) error {
	return f.UpdatePrimaryPhoneNumberFn(ctx, id, phoneNumber)
}

// UpdatePrimaryEmailAddress ...
func (f *FakeOnboardingRepository) UpdatePrimaryEmailAddress(ctx context.Context, id string, emailAddress string) error {
	return f.UpdatePrimaryEmailAddressFn(ctx, id, emailAddress)
}

// UpdateSecondaryPhoneNumbers ...
func (f *FakeOnboardingRepository) UpdateSecondaryPhoneNumbers(ctx context.Context, id string, phoneNumbers []string) error {
	return f.UpdateSecondaryPhoneNumbersFn(ctx, id, phoneNumbers)
}

// UpdateSecondaryEmailAddresses ...
func (f *FakeOnboardingRepository) UpdateSecondaryEmailAddresses(ctx context.Context, id string, emailAddresses []string) error {
	return f.UpdateSecondaryEmailAddressesFn(ctx, id, emailAddresses)
}

// UpdateSuspended ...
func (f *FakeOnboardingRepository) UpdateSuspended(ctx context.Context, id string, status bool) error {
	return f.UpdateSuspendedFn(ctx, id, status)
}

// UpdatePhotoUploadID ...
func (f *FakeOnboardingRepository) UpdatePhotoUploadID(ctx context.Context, id string, uploadID string) error {
	return f.UpdatePhotoUploadIDFn(ctx, id, uploadID)
}

// UpdateCovers ...
func (f *FakeOnboardingRepository) UpdateCovers(ctx context.Context, id string, covers []base.Cover) error {
	return f.UpdateCoversFn(ctx, id, covers)
}

// UpdatePushTokens ...
func (f *FakeOnboardingRepository) UpdatePushTokens(ctx context.Context, id string, pushToken []string) error {
	return f.UpdatePushTokensFn(ctx, id, pushToken)
}

// UpdatePermissions ...
func (f *FakeOnboardingRepository) UpdatePermissions(ctx context.Context, id string, perms []base.PermissionType) error {
	return f.UpdatePermissionsFn(ctx, id, perms)
}

// UpdateBioData ...
func (f *FakeOnboardingRepository) UpdateBioData(ctx context.Context, id string, data base.BioData) error {
	return f.UpdateBioDataFn(ctx, id, data)
}

// UpdateVerifiedIdentifiers ...
func (f *FakeOnboardingRepository) UpdateVerifiedIdentifiers(ctx context.Context, id string, identifiers []base.VerifiedIdentifier) error {
	return f.UpdateVerifiedIdentifiersFn(ctx, id, identifiers)
}

// UpdateVerifiedUIDS ...
func (f *FakeOnboardingRepository) UpdateVerifiedUIDS(ctx context.Context, id string, uids []string) error {
	return f.UpdateVerifiedUIDSFn(ctx, id, uids)
}

// GetOrCreatePhoneNumberUser ...
func (f *FakeOnboardingRepository) GetOrCreatePhoneNumberUser(ctx context.Context,
	phone string,
) (*resources.CreatedUserResponse, error) {
	return f.GetOrCreatePhoneNumberUserFn(ctx, phone)
}

// HardResetSecondaryPhoneNumbers ...
func (f *FakeOnboardingRepository) HardResetSecondaryPhoneNumbers(ctx context.Context, id string, phoneNumbers []string) error {
	return f.HardResetSecondaryPhoneNumbersFn(ctx, id, phoneNumbers)
}

// HardResetSecondaryEmailAddress ...
func (f *FakeOnboardingRepository) HardResetSecondaryEmailAddress(ctx context.Context, id string, newSecondaryEmails []string) error {
	return f.HardResetSecondaryPhoneNumbersFn(ctx, id, newSecondaryEmails)
}

// GetUserProfileAttributes ...
func (f *FakeOnboardingRepository) GetUserProfileAttributes(
	ctx context.Context,
	UIDs []string,
	attribute string,
) (map[string][]string, error) {
	return f.GetUserProfileAttributesFn(
		ctx,
		UIDs,
		attribute,
	)
}

// CheckIfExperimentParticipant ...
func (f *FakeOnboardingRepository) CheckIfExperimentParticipant(ctx context.Context, profileID string) (bool, error) {
	return f.CheckIfExperimentParticipantFn(ctx, profileID)
}

// AddUserAsExperimentParticipant ...
func (f *FakeOnboardingRepository) AddUserAsExperimentParticipant(ctx context.Context, profile *base.UserProfile) (bool, error) {
	return f.AddUserAsExperimentParticipantFn(ctx, profile)
}

// RemoveUserAsExperimentParticipant ...
func (f *FakeOnboardingRepository) RemoveUserAsExperimentParticipant(ctx context.Context, profile *base.UserProfile) (bool, error) {
	return f.RemoveUserAsExperimentParticipantFn(ctx, profile)
}

// UpdateAddresses ...
func (f *FakeOnboardingRepository) UpdateAddresses(
	ctx context.Context,
	id string,
	address base.Address,
	addressType base.AddressType,
) error {
	return f.UpdateAddressesFn(ctx, id, address, addressType)
}

// AddNHIFDetails ...
func (f *FakeOnboardingRepository) AddNHIFDetails(
	ctx context.Context,
	input resources.NHIFDetailsInput,
	profileID string,
) (*domain.NHIFDetails, error) {
	return f.AddNHIFDetailsFn(ctx, input, profileID)
}

// GetNHIFDetailsByProfileID ...
func (f *FakeOnboardingRepository) GetNHIFDetailsByProfileID(
	ctx context.Context,
	profileID string,
) (*domain.NHIFDetails, error) {
	return f.GetNHIFDetailsByProfileIDFn(ctx, profileID)
}

// GetUserCommunicationsSettings ...
func (f *FakeOnboardingRepository) GetUserCommunicationsSettings(ctx context.Context, profileID string) (*base.UserCommunicationsSetting, error) {
	return f.GetUserCommunicationsSettingsFn(ctx, profileID)
}

// SetUserCommunicationsSettings ...
func (f *FakeOnboardingRepository) SetUserCommunicationsSettings(ctx context.Context, profileID string,
	allowWhatsApp *bool, allowTextSms *bool, allowPush *bool, allowEmail *bool) (*base.UserCommunicationsSetting, error) {
	return f.SetUserCommunicationsSettingsFn(ctx, profileID, allowWhatsApp, allowTextSms, allowPush, allowEmail)
}

// UpdateCustomerProfile ...
func (f *FakeOnboardingRepository) UpdateCustomerProfile(
	ctx context.Context,
	profileID string,
	customer base.Customer,
) (*base.Customer, error) {
	return f.UpdateCustomerProfileFn(ctx, profileID, customer)
}