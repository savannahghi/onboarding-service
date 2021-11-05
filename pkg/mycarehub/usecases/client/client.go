package client

import (
	"context"
	"fmt"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure"
)

// IRegisterClient ...
type IRegisterClient interface {
	// TODO: the input client profile must not have an ID set
	//		validate identifiers when creating
	//		if the enrolment date is not supplied, set it automatically
	//		default to the client profile being active right after creation
	//		create a patient on FHIR (HealthRecordID
	//		if identifiers not supplied (e.g patient being created on app), set
	//			an internal identifier as the default. It should be updated later
	//			with the CCC number or other final identifier
	// TODO: ensure the user exists...supplied user ID
	// TODO: only register clients who've been counselled
	// TODO: consider: after successful registration, send invite link automatically
	RegisterClient(
		ctx context.Context,
		userInput *dto.UserInput,
		clientInput *dto.ClientProfileInput,
	) (*domain.ClientUserProfile, error)
}

// IAddClientIdentifier ...
type IAddClientIdentifier interface {
	// TODO use idType and settings to decide if it's a primary identifier or not
	//AddIdentifier(ctx context.Context, clientID string, idType enums.IdentifierType, idValue string, isPrimary bool) (*domain.Identifier, error)
}

// IInactivateClient ...
type IInactivateClient interface {
	// TODO Consider making reasons an enum
	//InactivateClient(clientID string, reason string, notes string) (bool, error)
}

// IReactivateClient ...
type IReactivateClient interface {
	//ReactivateClient(clientID string, reason string, notes string) (bool, error)
}

// ITransferClient ...
type ITransferClient interface {
	// TODO: maintain log of past transfers, who did it etc
	// TransferClient(
	// 	ctx context.Context,
	// 	clientID string,
	// 	originFacilityID string,
	// 	destinationFacilityID string,
	// 	reason enums.TransferReason,
	// 	notes string,
	// ) (bool, error)
}

// IGetClientIdentifiers ...
type IGetClientIdentifiers interface {
	//GetIdentifiers(clientID string, active bool) ([]*domain.Identifier, error)
}

// IInactivateClientIdentifier ...
type IInactivateClientIdentifier interface {
	//InactivateIdentifier(clientID string, identifierID string) (bool, error)
}

// IAssignTreatmentSupporter ...
type IAssignTreatmentSupporter interface {
	// AssignTreatmentSupporter(
	// 	clientID string,
	// 	treatmentSupporterID string,
	// 	treatmentSupporterType string, // TODO: enum, start with CHV and Treatment buddy
	// ) (bool, error)
}

// IUnassignTreatmentSupporter ...
type IUnassignTreatmentSupporter interface {
	// UnassignTreatmentSupporter(
	// 	clientID string,
	// 	treatmentSupporterID string,
	// 	reason string, // TODO: ensure these are in an audit log
	// 	notes string, // TODO: Optional
	// ) (bool, error)
}

// IAddRelatedPerson ...
type IAddRelatedPerson interface {
	// add next of kin
	// 	AddRelatedPerson(
	// 		clientID string,
	// 		relatedPerson *domain.RelatedPerson,
	// 	) (*domain.RelatedPerson, bool)
}

// UseCasesClientProfile ...
type UseCasesClientProfile interface {
	IAddClientIdentifier
	IGetClientIdentifiers
	IInactivateClientIdentifier
	IRegisterClient
	IInactivateClient
	IReactivateClient
	ITransferClient
	IAssignTreatmentSupporter
	IUnassignTreatmentSupporter
	IAddRelatedPerson
}

// UseCasesClientImpl represents user implementation object
type UseCasesClientImpl struct {
	Create infrastructure.Create
	Query  infrastructure.Query
	Delete infrastructure.Delete
}

// NewUseCasesClientImpl returns a new Client service
func NewUseCasesClientImpl(create infrastructure.Create, query infrastructure.Query, delete infrastructure.Delete) *UseCasesClientImpl {
	return &UseCasesClientImpl{
		Create: create,
		Query:  query,
		Delete: delete,
	}
}

// RegisterClient registers a client into the platform
func (cl *UseCasesClientImpl) RegisterClient(
	ctx context.Context,
	userInput *dto.UserInput,
	clientInput *dto.ClientProfileInput,
) (*domain.ClientUserProfile, error) {

	if userInput == nil {
		return nil, fmt.Errorf("user input cannot be nil")
	}

	if clientInput == nil {
		return nil, fmt.Errorf("client input cannot be nil")
	}

	return cl.Create.RegisterClient(ctx, userInput, clientInput)
}

// AddIdentifier stages and adds client identifiers
func (cl *UseCasesClientImpl) AddIdentifier(ctx context.Context, clientID string, idType enums.IdentifierType, idValue string, isPrimary bool) (*domain.Identifier, error) {
	return nil, nil
}

// InactivateClient makes a client inactive and removes the client from the list of active users
func (cl *UseCasesClientImpl) InactivateClient(clientID string, reason string, notes string) (bool, error) {
	return true, nil
}

// ReactivateClient makes inactive client active and returns the client to the list of active user
func (cl *UseCasesClientImpl) ReactivateClient(clientID string, reason string, notes string) (bool, error) {
	return true, nil
}

// TransferClient transfers a client from one facility to another facility for various reasons
func (cl *UseCasesClientImpl) TransferClient(
	ctx context.Context,
	clientID string,
	originFacilityID string,
	destinationFacilityID string,
	reason enums.TransferReason,
	notes string,
) (bool, error) {
	return false, nil
}

// GetIdentifiers fetches and returns a list of client active identifiers
func (cl *UseCasesClientImpl) GetIdentifiers(clientID string, active bool) ([]*domain.Identifier, error) {
	return nil, nil
}

// InactivateIdentifier toggles and make client identifier inactive
func (cl *UseCasesClientImpl) InactivateIdentifier(clientID string, identifierID string) (bool, error) {
	return true, nil
}

// AssignTreatmentSupporter assigns a treatment supporter to a client
func (cl *UseCasesClientImpl) AssignTreatmentSupporter(
	clientID string,
	treatmentSupporterID string,
	treatmentSupporterType string, // TODO: enum, start with CHV and Treatment buddy
) (bool, error) {
	return true, nil
}

// UnassignTreatmentSupporter unassign treatment supporter from a client
func (cl *UseCasesClientImpl) UnassignTreatmentSupporter(
	clientID string,
	treatmentSupporterID string,
	reason string, // TODO: ensure these are in an audit log
	notes string, // TODO: Optional
) (bool, error) {
	return true, nil
}

// AddRelatedPerson adds client related person. The related person here is like Next of Kin
func (cl *UseCasesClientImpl) AddRelatedPerson(
	clientID string,
	relatedPerson *domain.RelatedPerson,
) (*domain.RelatedPerson, bool) {
	return nil, false
}