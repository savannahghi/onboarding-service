package servicerequest

import (
	"context"
	"fmt"

	"github.com/savannahghi/feedlib"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/common/helpers"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/exceptions"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/extension"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/usecases/user"
)

// Service requests are tasks for the healthcare staff on the platform. Some examples are:
// red flags raised by content posted to a group or entries into health diaries
// appointment reschedule requests
// These tasks will be presented on a list and notified (e.g., via push notifications).
// Each task will have a status. When first created, the tasks will be marked as “PENDING”.
// Once the relevant actions are taken, it will be possible to mark them as “IN PROGRESS”, “RESOLVED” and add relevant notes.
// In order to ensure that a task is not addressed by multiple people at the same time, each task will be updated with a record of the user and timestamp each time the status is changed.

// ICreateServiceRequest is an interface that holds the method signature for creating a service request
type ICreateServiceRequest interface {
	CreateServiceRequest(
		ctx context.Context,
		clientID string,
		requestType, request, cccNumber string,
	) (bool, error)

	CreatePinResetServiceRequest(ctx context.Context, cccNumber string) (bool, error)
}

// ISetInProgresssBy is an interface that contains the method signature for assigning the staff currently working on a request
type ISetInProgresssBy interface {
	SetInProgressBy(ctx context.Context, requestID string, staffID string) (bool, error)
}

// IGetServiceRequests is an interface that holds the method signature for getting service requests
type IGetServiceRequests interface {
	GetServiceRequests(ctx context.Context, requestType, requestStatus, facilityID *string) ([]*domain.ServiceRequest, error)
	GetServiceRequestsForKenyaEMR(ctx context.Context, payload *dto.ServiceRequestPayload) ([]*domain.ServiceRequest, error)
	GetPendingServiceRequestsCount(ctx context.Context, facilityID string) (*domain.ServiceRequestsCount, error)
}

// IResolveServiceRequest is an interface that holds the method signature for resolving a service request
type IResolveServiceRequest interface {
	ResolveServiceRequest(ctx context.Context, staffID *string, serviceRequestID *string) (bool, error)
	ApprovePinResetServiceRequest(
		ctx context.Context,
		clientID string,
		serviceRequestID string,
		cccNumber string,
		phoneNumber string,
		physicalIdentityVerified bool,
	) (bool, error)
}

// IUpdateServiceRequest is the interface holding the method signature for updating service requests.
type IUpdateServiceRequest interface {
	UpdateServiceRequestsFromKenyaEMR(ctx context.Context, payload *dto.UpdateServiceRequestsPayload) (bool, error)
}

// UseCaseServiceRequest holds all the interfaces that represent the service request business logic
type UseCaseServiceRequest interface {
	ICreateServiceRequest
	IGetServiceRequests
	ISetInProgresssBy
	IResolveServiceRequest
	IUpdateServiceRequest
}

// UseCasesServiceRequestImpl embeds the service request logic
type UseCasesServiceRequestImpl struct {
	Create      infrastructure.Create
	Query       infrastructure.Query
	Update      infrastructure.Update
	ExternalExt extension.ExternalMethodsExtension
	User        user.UseCasesUser
}

// NewUseCaseServiceRequestImpl creates a new service request instance
func NewUseCaseServiceRequestImpl(
	create infrastructure.Create,
	query infrastructure.Query,
	update infrastructure.Update,
	ext extension.ExternalMethodsExtension,
	user user.UseCasesUser,
) *UseCasesServiceRequestImpl {
	return &UseCasesServiceRequestImpl{
		Create:      create,
		Query:       query,
		Update:      update,
		ExternalExt: ext,
		User:        user,
	}
}

// CreateServiceRequest creates a service request
func (u *UseCasesServiceRequestImpl) CreateServiceRequest(
	ctx context.Context,
	clientID string,
	requestType, request, cccNumber string,
) (bool, error) {
	clientProfile, err := u.Query.GetClientProfileByClientID(ctx, clientID)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.ClientProfileNotFoundErr(err)
	}
	serviceRequest := &domain.ClientServiceRequest{
		Active:      true,
		RequestType: requestType,
		Request:     request,
		Status:      "PENDING",
		ClientID:    clientID,
		FacilityID:  clientProfile.FacilityID,
		CCCNumber:   cccNumber,
	}
	err = u.Create.CreateServiceRequest(ctx, serviceRequest)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to create service request: %v", err)
	}
	return true, nil
}

// SetInProgressBy assigns to a service request, staff currently working on the service request
func (u *UseCasesServiceRequestImpl) SetInProgressBy(ctx context.Context, requestID string, staffID string) (bool, error) {
	if requestID == "" || staffID == "" {
		return false, fmt.Errorf("request ID or staff ID cannot be empty")
	}
	return u.Update.SetInProgressBy(ctx, requestID, staffID)
}

// GetServiceRequests gets service requests based on the parameters provided
func (u *UseCasesServiceRequestImpl) GetServiceRequests(
	ctx context.Context,
	requestType *string,
	requestStatus *string,
	facilityID *string,
) ([]*domain.ServiceRequest, error) {
	if requestType != nil {
		if !enums.ServiceRequestType(*requestType).IsValid() {
			return nil, fmt.Errorf("invalid request type: %v", *requestType)
		}
	}
	if requestStatus != nil {
		if !enums.ServiceRequestStatus(*requestStatus).IsValid() {
			return nil, fmt.Errorf("invalid request status: %v", *requestStatus)
		}
	}

	return u.Query.GetServiceRequests(ctx, requestType, requestStatus, facilityID)
}

// GetServiceRequestsForKenyaEMR fetches all the most recent service requests  that have not been
// synced to KenyaEMR.
func (u *UseCasesServiceRequestImpl) GetServiceRequestsForKenyaEMR(ctx context.Context, payload *dto.ServiceRequestPayload) ([]*domain.ServiceRequest, error) {
	return u.Query.GetServiceRequestsForKenyaEMR(ctx, payload)
}

// GetPendingServiceRequestsCount gets the total number of service requests
func (u *UseCasesServiceRequestImpl) GetPendingServiceRequestsCount(ctx context.Context, facilityID string) (*domain.ServiceRequestsCount, error) {
	if facilityID == "" {
		return nil, fmt.Errorf("facility id cannot be empty")
	}

	return u.Query.GetPendingServiceRequestsCount(ctx, facilityID)
}

// ResolveServiceRequest resolves a service request
func (u *UseCasesServiceRequestImpl) ResolveServiceRequest(ctx context.Context, staffID *string, serviceRequestID *string) (bool, error) {
	if staffID == nil {
		return false, fmt.Errorf("staff ID is required")
	}
	if serviceRequestID == nil {
		return false, fmt.Errorf("service request ID is required")
	}
	ok, err := u.Update.ResolveServiceRequest(ctx, staffID, serviceRequestID)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to update service request: %v", err)
	}

	if !ok {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to resolve service request")
	}

	return ok, nil
}

// UpdateServiceRequestsFromKenyaEMR is used to update service requests from KenyaEMR to MyCareHub service.
func (u *UseCasesServiceRequestImpl) UpdateServiceRequestsFromKenyaEMR(ctx context.Context, payload *dto.UpdateServiceRequestsPayload) (bool, error) {
	var serviceRequests []domain.ServiceRequest
	for _, v := range payload.ServiceRequests {
		serviceRequest := &domain.ServiceRequest{
			ID:           v.ID,
			RequestType:  v.RequestType,
			Status:       v.Status,
			InProgressAt: &v.InProgressAt,
			InProgressBy: &v.InProgressBy,
			ResolvedAt:   &v.ResolvedAt,
			ResolvedBy:   &v.ResolvedBy,
		}

		serviceRequests = append(serviceRequests, *serviceRequest)
	}

	serviceReq := &domain.UpdateServiceRequestsPayload{
		ServiceRequests: serviceRequests,
	}
	return u.Update.UpdateServiceRequests(ctx, serviceReq)
}

// CreatePinResetServiceRequest creates a PIN_RESET service request. This occurs when a user attempts to change
// their pin but they don't succeed.
func (u *UseCasesServiceRequestImpl) CreatePinResetServiceRequest(ctx context.Context, cccNumber string) (bool, error) {
	// TODO: Check if the service request exists before creating a new one
	clientProfile, err := u.Query.GetClientProfileByCCCNumber(ctx, cccNumber)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.GetError(err)
	}

	request := "Request to change pin"
	_, err = u.CreateServiceRequest(
		ctx,
		*clientProfile.ID,
		string(enums.ServiceRequestTypePinReset),
		request,
		cccNumber,
	)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.InternalErr(err)
	}

	return true, nil
}

// ApprovePinResetServiceRequest is used to approve a pin reset service request. This is used by the
// healthcare worker to reset the login credentials of a user who failed to login and requested for help from
// the health care worker.
//
// The basic workflow is
// 1. Get the logged in user ID - This will be used to identify the staff who resolved the request
// 2. Verify that the patient was physically verified by the healthcare worker and that the provided
// ccc number matches the one on their profile
// 3. Mark the service request as IN_PROGRESS
// 4. Send a fresh invite to the user and invalidate the previous pins
// 5. Update the field `pin_change_required` to true and mark the service request as resolved
func (u *UseCasesServiceRequestImpl) ApprovePinResetServiceRequest(
	ctx context.Context,
	clientID string,
	serviceRequestID string,
	cccNumber string,
	phoneNumber string,
	physicalIdentityVerified bool,
) (bool, error) {
	flavour := feedlib.FlavourConsumer
	loggedInUserID, err := u.ExternalExt.GetLoggedInUserUID(ctx)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.GetLoggedInUserUIDErr(err)
	}

	staff, err := u.Query.GetStaffProfileByUserID(ctx, loggedInUserID)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, exceptions.StaffProfileNotFoundErr(err)
	}

	if !physicalIdentityVerified {
		return false, fmt.Errorf("the patient has not been physically verified by the healthcare worker")
	}

	identifier, err := u.Query.GetClientCCCIdentifier(ctx, clientID)
	if err != nil {
		return false, err
	}

	if identifier == nil {
		return false, fmt.Errorf("patient has no recorded identifier")
	}

	if cccNumber != identifier.IdentifierValue {
		return false, fmt.Errorf("the ccc number provided does not match with the one on the patient profile")
	}

	_, err = u.SetInProgressBy(ctx, serviceRequestID, *staff.ID)
	if err != nil {
		return false, err
	}

	user, err := u.Query.GetUserProfileByPhoneNumber(ctx, phoneNumber, flavour)
	if err != nil {
		return false, err
	}

	_, err = u.User.InviteUser(ctx, *user.ID, phoneNumber, flavour)
	if err != nil {
		return false, err
	}

	err = u.Update.UpdateUserPinChangeRequiredStatus(ctx, *user.ID, flavour, true)
	if err != nil {
		return false, err
	}

	_, err = u.ResolveServiceRequest(ctx, staff.ID, &serviceRequestID)
	if err != nil {
		return false, err
	}

	return true, nil
}
