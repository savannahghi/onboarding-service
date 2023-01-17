package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/savannahghi/feedlib"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

// AcceptTerms is the resolver for the acceptTerms field.
func (r *mutationResolver) AcceptTerms(ctx context.Context, userID string, termsID int) (bool, error) {
	r.checkPreconditions()
	return r.mycarehub.Terms.AcceptTerms(ctx, &userID, &termsID)
}

// SetNickName is the resolver for the setNickName field.
func (r *mutationResolver) SetNickName(ctx context.Context, userID string, nickname string) (bool, error) {
	r.checkPreconditions()
	return r.mycarehub.User.SetNickName(ctx, userID, nickname)
}

// CompleteOnboardingTour is the resolver for the completeOnboardingTour field.
func (r *mutationResolver) CompleteOnboardingTour(ctx context.Context, userID string, flavour feedlib.Flavour) (bool, error) {
	r.checkPreconditions()
	return r.mycarehub.User.CompleteOnboardingTour(ctx, userID, flavour)
}

// RegisterClient is the resolver for the registerClient field.
func (r *mutationResolver) RegisterClient(ctx context.Context, input *dto.ClientRegistrationInput) (*dto.ClientRegistrationOutput, error) {
	return r.mycarehub.User.RegisterClient(ctx, input)
}

// RegisterStaff is the resolver for the registerStaff field.
func (r *mutationResolver) RegisterStaff(ctx context.Context, input dto.StaffRegistrationInput) (*dto.StaffRegistrationOutput, error) {
	return r.mycarehub.User.RegisterStaff(ctx, input)
}

// RegisterCaregiver is the resolver for the registerCaregiver field.
func (r *mutationResolver) RegisterCaregiver(ctx context.Context, input dto.CaregiverInput) (*domain.CaregiverProfile, error) {
	return r.mycarehub.User.RegisterCaregiver(ctx, input)
}

// RegisterClientAsCaregiver is the resolver for the registerClientAsCaregiver field.
func (r *mutationResolver) RegisterClientAsCaregiver(ctx context.Context, clientID string, caregiverNumber string) (*domain.CaregiverProfile, error) {
	return r.mycarehub.User.RegisterClientAsCaregiver(ctx, clientID, caregiverNumber)
}

// OptOut is the resolver for the optOut field.
func (r *mutationResolver) OptOut(ctx context.Context, phoneNumber string, flavour feedlib.Flavour) (bool, error) {
	return r.mycarehub.User.Consent(ctx, phoneNumber, flavour)
}

// SetPushToken is the resolver for the setPushToken field.
func (r *mutationResolver) SetPushToken(ctx context.Context, token string) (bool, error) {
	return r.mycarehub.User.RegisterPushToken(ctx, token)
}

// InviteUser is the resolver for the inviteUser field.
func (r *mutationResolver) InviteUser(ctx context.Context, userID string, phoneNumber string, flavour feedlib.Flavour, reinvite *bool) (bool, error) {
	return r.mycarehub.User.InviteUser(ctx, userID, phoneNumber, flavour, *reinvite)
}

// SetUserPin is the resolver for the setUserPIN field.
func (r *mutationResolver) SetUserPin(ctx context.Context, input *dto.PINInput) (bool, error) {
	return r.mycarehub.User.SetUserPIN(ctx, *input)
}

// TransferClientToFacility is the resolver for the transferClientToFacility field.
func (r *mutationResolver) TransferClientToFacility(ctx context.Context, clientID string, facilityID string) (bool, error) {
	return r.mycarehub.User.TransferClientToFacility(ctx, &clientID, &facilityID)
}

// SetStaffDefaultFacility is the resolver for the setStaffDefaultFacility field.
func (r *mutationResolver) SetStaffDefaultFacility(ctx context.Context, staffID string, facilityID string) (*domain.Facility, error) {
	return r.mycarehub.User.SetStaffDefaultFacility(ctx, staffID, facilityID)
}

// SetClientDefaultFacility is the resolver for the setClientDefaultFacility field.
func (r *mutationResolver) SetClientDefaultFacility(ctx context.Context, clientID string, facilityID string) (*domain.Facility, error) {
	return r.mycarehub.User.SetClientDefaultFacility(ctx, clientID, facilityID)
}

// AddFacilitiesToStaffProfile is the resolver for the addFacilitiesToStaffProfile field.
func (r *mutationResolver) AddFacilitiesToStaffProfile(ctx context.Context, staffID string, facilities []string) (bool, error) {
	return r.mycarehub.User.AddFacilitiesToStaffProfile(ctx, staffID, facilities)
}

// AddFacilitiesToClientProfile is the resolver for the addFacilitiesToClientProfile field.
func (r *mutationResolver) AddFacilitiesToClientProfile(ctx context.Context, clientID string, facilities []string) (bool, error) {
	return r.mycarehub.User.AddFacilitiesToClientProfile(ctx, clientID, facilities)
}

// RemoveFacilitiesFromClientProfile is the resolver for the removeFacilitiesFromClientProfile field.
func (r *mutationResolver) RemoveFacilitiesFromClientProfile(ctx context.Context, clientID string, facilities []string) (bool, error) {
	return r.mycarehub.User.RemoveFacilitiesFromClientProfile(ctx, clientID, facilities)
}

// AssignCaregiver is the resolver for the assignCaregiver field.
func (r *mutationResolver) AssignCaregiver(ctx context.Context, input dto.ClientCaregiverInput) (bool, error) {
	return r.mycarehub.User.AssignCaregiver(ctx, input)
}

// RemoveFacilitiesFromStaffProfile is the resolver for the removeFacilitiesFromStaffProfile field.
func (r *mutationResolver) RemoveFacilitiesFromStaffProfile(ctx context.Context, staffID string, facilities []string) (bool, error) {
	return r.mycarehub.User.RemoveFacilitiesFromStaffProfile(ctx, staffID, facilities)
}

// RegisterExistingUserAsStaff is the resolver for the registerExistingUserAsStaff field.
func (r *mutationResolver) RegisterExistingUserAsStaff(ctx context.Context, input dto.ExistingUserStaffInput) (*dto.StaffRegistrationOutput, error) {
	r.checkPreconditions()

	return r.mycarehub.User.RegisterExistingUserAsStaff(ctx, input)
}

// ConsentToAClientCaregiver is the resolver for the consentToAClientCaregiver field.
func (r *mutationResolver) ConsentToAClientCaregiver(ctx context.Context, clientID string, caregiverID string, consent bool) (bool, error) {
	return r.mycarehub.User.ConsentToAClientCaregiver(ctx, clientID, caregiverID, consent)
}

// ConsentToManagingClient is the resolver for the consentToManagingClient field.
func (r *mutationResolver) ConsentToManagingClient(ctx context.Context, caregiverID string, clientID string, consent bool) (bool, error) {
	return r.mycarehub.User.ConsentToManagingClient(ctx, caregiverID, clientID, consent)
}

// RegisterExistingUserAsClient is the resolver for the registerExistingUserAsClient field.
func (r *mutationResolver) RegisterExistingUserAsClient(ctx context.Context, input dto.ExistingUserClientInput) (*dto.ClientRegistrationOutput, error) {
	r.checkPreconditions()

	return r.mycarehub.User.RegisterExistingUserAsClient(ctx, input)
}

// SetCaregiverCurrentClient is the resolver for the setCaregiverCurrentClient field.
func (r *mutationResolver) SetCaregiverCurrentClient(ctx context.Context, clientID string) (*domain.ClientProfile, error) {
	return r.mycarehub.User.SetCaregiverCurrentClient(ctx, clientID)
}

// SetCaregiverCurrentFacility is the resolver for the setCaregiverCurrentFacility field.
func (r *mutationResolver) SetCaregiverCurrentFacility(ctx context.Context, clientID string, facilityID string) (*domain.Facility, error) {
	return r.mycarehub.User.SetCaregiverCurrentFacility(ctx, clientID, facilityID)
}

// RegisterExistingUserAsCaregiver is the resolver for the registerExistingUserAsCaregiver field.
func (r *mutationResolver) RegisterExistingUserAsCaregiver(ctx context.Context, userID string, caregiverNumber string) (*domain.CaregiverProfile, error) {
	r.checkPreconditions()

	return r.mycarehub.User.RegisterExistingUserAsCaregiver(ctx, userID, caregiverNumber)
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, userID string, cccNumber *string, username *string, phoneNumber *string, programID string, flavour feedlib.Flavour) (bool, error) {
	r.checkPreconditions()

	return r.mycarehub.User.UpdateUserProfile(ctx, userID, cccNumber, username, phoneNumber, programID, flavour)
}

// GetCurrentTerms is the resolver for the getCurrentTerms field.
func (r *queryResolver) GetCurrentTerms(ctx context.Context) (*domain.TermsOfService, error) {
	r.checkPreconditions()
	return r.mycarehub.Terms.GetCurrentTerms(ctx)
}

// VerifyPin is the resolver for the verifyPIN field.
func (r *queryResolver) VerifyPin(ctx context.Context, userID string, flavour feedlib.Flavour, pin string) (bool, error) {
	return r.mycarehub.User.VerifyPIN(ctx, userID, flavour, pin)
}

// SearchClientUser is the resolver for the searchClientUser field.
func (r *queryResolver) SearchClientUser(ctx context.Context, searchParameter string) ([]*domain.ClientProfile, error) {
	return r.mycarehub.User.SearchClientUser(ctx, searchParameter)
}

// SearchStaffUser is the resolver for the searchStaffUser field.
func (r *queryResolver) SearchStaffUser(ctx context.Context, searchParameter string) ([]*domain.StaffProfile, error) {
	return r.mycarehub.User.SearchStaffUser(ctx, searchParameter)
}

// SearchCaregiverUser is the resolver for the searchCaregiverUser field.
func (r *queryResolver) SearchCaregiverUser(ctx context.Context, searchParameter string) ([]*domain.CaregiverProfile, error) {
	return r.mycarehub.User.SearchCaregiverUser(ctx, searchParameter)
}

// GetClientProfileByCCCNumber is the resolver for the getClientProfileByCCCNumber field.
func (r *queryResolver) GetClientProfileByCCCNumber(ctx context.Context, cCCNumber string) (*domain.ClientProfile, error) {
	return r.mycarehub.User.GetClientProfileByCCCNumber(ctx, cCCNumber)
}

// GetUserLinkedFacilities is the resolver for the getUserLinkedFacilities field.
func (r *queryResolver) GetUserLinkedFacilities(ctx context.Context, userID string, paginationInput dto.PaginationsInput) (*dto.FacilityOutputPage, error) {
	return r.mycarehub.User.GetUserLinkedFacilities(ctx, userID, paginationInput)
}

// GetCaregiverManagedClients is the resolver for the getCaregiverManagedClients field.
func (r *queryResolver) GetCaregiverManagedClients(ctx context.Context, userID string, paginationInput dto.PaginationsInput) (*dto.ManagedClientOutputPage, error) {
	return r.mycarehub.User.GetCaregiverManagedClients(ctx, userID, paginationInput)
}

// ListClientsCaregivers is the resolver for the listClientsCaregivers field.
func (r *queryResolver) ListClientsCaregivers(ctx context.Context, clientID string, paginationInput *dto.PaginationsInput) (*dto.CaregiverProfileOutputPage, error) {
	return r.mycarehub.User.ListClientsCaregivers(ctx, clientID, paginationInput)
}

// GetStaffFacilities is the resolver for the getStaffFacilities field.
func (r *queryResolver) GetStaffFacilities(ctx context.Context, staffID string, paginationInput dto.PaginationsInput) (*dto.FacilityOutputPage, error) {
	r.checkPreconditions()

	return r.mycarehub.User.GetStaffFacilities(ctx, staffID, paginationInput)
}

// GetClientFacilities is the resolver for the getClientFacilities field.
func (r *queryResolver) GetClientFacilities(ctx context.Context, clientID string, paginationInput dto.PaginationsInput) (*dto.FacilityOutputPage, error) {
	r.checkPreconditions()

	return r.mycarehub.User.GetClientFacilities(ctx, clientID, paginationInput)
}
