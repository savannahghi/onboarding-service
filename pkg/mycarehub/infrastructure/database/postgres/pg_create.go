package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/savannahghi/enumutils"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/common/helpers"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure/database/postgres/gorm"
)

// GetOrCreateFacility is responsible from creating a representation of a facility
// A facility here is the healthcare facility that are on the platform.
// A facility MFL CODE must be unique across the platform. I forms part of the unique identifiers
//
// TODO: Create a helper the checks for all required fields
// TODO: Make the create method idempotent
func (d *MyCareHubDb) GetOrCreateFacility(ctx context.Context, facility *dto.FacilityInput) (*domain.Facility, error) {
	if err := facility.Validate(); err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, fmt.Errorf("facility input validation failed: %s", err)
	}

	facilityObj := &gorm.Facility{
		Name:               facility.Name,
		Code:               facility.Code,
		Active:             facility.Active,
		County:             facility.County,
		Phone:              facility.Phone,
		Description:        facility.Description,
		FHIROrganisationID: facility.FHIROrganisationID,
	}

	facilitySession, err := d.create.GetOrCreateFacility(ctx, facilityObj)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, fmt.Errorf("failed to create facility: %v", err)
	}

	return d.mapFacilityObjectToDomain(facilitySession), nil
}

// SaveTemporaryUserPin does the actual saving of the users PIN in the database
func (d *MyCareHubDb) SaveTemporaryUserPin(ctx context.Context, pinData *domain.UserPIN) (bool, error) {
	pinObj := &gorm.PINData{
		UserID:    pinData.UserID,
		HashedPIN: pinData.HashedPIN,
		ValidFrom: pinData.ValidFrom,
		ValidTo:   pinData.ValidTo,
		IsValid:   pinData.IsValid,
		Flavour:   pinData.Flavour,
		Salt:      pinData.Salt,
	}

	_, err := d.create.SaveTemporaryUserPin(ctx, pinObj)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to save user pin: %v", err)
	}

	return true, nil
}

// SavePin gets the pin details from the user and saves it in the database
func (d *MyCareHubDb) SavePin(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {

	pinObj := &gorm.PINData{
		UserID:    pinInput.UserID,
		HashedPIN: pinInput.HashedPIN,
		ValidFrom: pinInput.ValidFrom,
		ValidTo:   pinInput.ValidTo,
		IsValid:   pinInput.IsValid,
		Flavour:   pinInput.Flavour,
		Salt:      pinInput.Salt,
	}

	_, err := d.create.SavePin(ctx, pinObj)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return false, fmt.Errorf("failed to save user pin: %v", err)
	}

	return true, nil
}

// SaveOTP saves the otp to the database
func (d *MyCareHubDb) SaveOTP(ctx context.Context, otpInput *domain.OTP) error {
	otpObject := &gorm.UserOTP{
		UserID:      otpInput.UserID,
		Valid:       otpInput.Valid,
		GeneratedAt: otpInput.GeneratedAt,
		ValidUntil:  otpInput.ValidUntil,
		Channel:     otpInput.Channel,
		PhoneNumber: otpInput.PhoneNumber,
		Flavour:     otpInput.Flavour,
		OTP:         otpInput.OTP,
	}

	err := d.create.SaveOTP(ctx, otpObject)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return fmt.Errorf("failed to save OTP")
	}

	return nil
}

// SaveSecurityQuestionResponse saves the security question response to the database
func (d *MyCareHubDb) SaveSecurityQuestionResponse(ctx context.Context, securityQuestionResponse []*dto.SecurityQuestionResponseInput) error {
	var securityQuestionResponseObj []*gorm.SecurityQuestionResponse
	for _, sqr := range securityQuestionResponse {
		response := &gorm.SecurityQuestionResponse{
			UserID:     sqr.UserID,
			QuestionID: sqr.SecurityQuestionID,
			Active:     true,
			Response:   sqr.Response,
		}
		securityQuestionResponseObj = append(securityQuestionResponseObj, response)
	}

	err := d.create.SaveSecurityQuestionResponse(ctx, securityQuestionResponseObj)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return fmt.Errorf("failed to save security question response data")
	}

	return nil
}

// CreateHealthDiaryEntry is used to add a health diary record to the database.
func (d *MyCareHubDb) CreateHealthDiaryEntry(ctx context.Context, healthDiaryInput *domain.ClientHealthDiaryEntry) error {
	healthDiaryResponse := &gorm.ClientHealthDiaryEntry{
		Active:                healthDiaryInput.Active,
		Mood:                  healthDiaryInput.Mood,
		Note:                  healthDiaryInput.Note,
		EntryType:             healthDiaryInput.EntryType,
		ShareWithHealthWorker: healthDiaryInput.ShareWithHealthWorker,
		SharedAt:              healthDiaryInput.SharedAt,
		ClientID:              healthDiaryInput.ClientID,
	}

	err := d.create.CreateHealthDiaryEntry(ctx, healthDiaryResponse)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return err
	}

	return nil
}

// CreateServiceRequest creates  a service request which will be handled by a staff user.
// This happens in a transaction because we do not want to
// create a health diary entry without a subsequent service request when the client's mood is "VERY_BAD"
func (d *MyCareHubDb) CreateServiceRequest(ctx context.Context, serviceRequestInput *dto.ServiceRequestInput) error {
	meta, err := json.Marshal(serviceRequestInput.Meta)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return fmt.Errorf("failed to marshal meta data: %v", err)
	}
	serviceRequest := &gorm.ClientServiceRequest{
		Active:      serviceRequestInput.Active,
		RequestType: serviceRequestInput.RequestType,
		Request:     serviceRequestInput.Request,
		Status:      serviceRequestInput.Status,
		ClientID:    serviceRequestInput.ClientID,
		FacilityID:  serviceRequestInput.FacilityID,
		Meta:        string(meta),
	}

	err = d.create.CreateServiceRequest(ctx, serviceRequest)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return err
	}

	return nil
}

// CreateStaffServiceRequest creates a new service request for the specified staff
func (d *MyCareHubDb) CreateStaffServiceRequest(ctx context.Context, serviceRequestInput *dto.ServiceRequestInput) error {
	meta, err := json.Marshal(serviceRequestInput.Meta)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return fmt.Errorf("failed to marshal meta data: %v", err)
	}
	serviceRequest := &gorm.StaffServiceRequest{
		Active:            serviceRequestInput.Active,
		RequestType:       serviceRequestInput.RequestType,
		Request:           serviceRequestInput.Request,
		Status:            serviceRequestInput.Status,
		StaffID:           serviceRequestInput.StaffID,
		DefaultFacilityID: &serviceRequestInput.FacilityID,
		Meta:              string(meta),
	}

	err = d.create.CreateStaffServiceRequest(ctx, serviceRequest)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return err
	}

	return nil
}

// CreateClientCaregiver creates a client's caregiver
func (d *MyCareHubDb) CreateClientCaregiver(ctx context.Context, caregiverInput *dto.CaregiverInput) error {
	caregiver := &gorm.Caregiver{
		FirstName:     caregiverInput.FirstName,
		LastName:      caregiverInput.LastName,
		PhoneNumber:   caregiverInput.PhoneNumber,
		CaregiverType: caregiverInput.CaregiverType,
	}

	err := d.create.CreateClientCaregiver(ctx, caregiverInput.ClientID, caregiver)
	if err != nil {
		return err
	}

	return nil
}

// CreateCommunity creates a channel in the database
func (d *MyCareHubDb) CreateCommunity(ctx context.Context, communityInput *dto.CommunityInput) (*domain.Community, error) {

	var genderList pq.StringArray
	for _, g := range communityInput.Gender {
		genderList = append(genderList, string(*g))
	}

	var clientTypeList pq.StringArray
	for _, c := range communityInput.ClientType {
		clientTypeList = append(clientTypeList, string(*c))
	}

	input := &gorm.Community{
		Name:         communityInput.Name,
		Description:  communityInput.Description,
		Active:       true,
		MinimumAge:   communityInput.AgeRange.LowerBound,
		MaximumAge:   communityInput.AgeRange.UpperBound,
		Gender:       genderList,
		ClientTypes:  clientTypeList,
		InviteOnly:   communityInput.InviteOnly,
		Discoverable: true,
	}

	channel, err := d.create.CreateCommunity(ctx, input)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return nil, err
	}

	var genders []enumutils.Gender
	for _, k := range channel.Gender {
		genders = append(genders, enumutils.Gender(k))
	}

	var clientTypes []enums.ClientType
	for _, k := range channel.ClientTypes {
		clientTypes = append(clientTypes, enums.ClientType(k))
	}

	return &domain.Community{
		ID:          channel.ID,
		Name:        channel.Name,
		Description: channel.Description,
		AgeRange: &domain.AgeRange{
			LowerBound: channel.MinimumAge,
			UpperBound: channel.MaximumAge,
		},
		Gender:     genders,
		ClientType: clientTypes,
		InviteOnly: channel.InviteOnly,
	}, nil
}

// GetOrCreateNextOfKin creates a related person who is a next of kin
func (d *MyCareHubDb) GetOrCreateNextOfKin(ctx context.Context, person *dto.NextOfKinPayload, clientID, contactID string) error {

	pn := &gorm.RelatedPerson{
		FirstName:        person.Name,
		RelationshipType: "NEXT_OF_KIN",
	}

	return d.create.GetOrCreateNextOfKin(ctx, pn, clientID, contactID)
}

// GetOrCreateContact creates a contact
func (d *MyCareHubDb) GetOrCreateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {

	ct := &gorm.Contact{
		Active:       true,
		ContactType:  contact.ContactType,
		ContactValue: contact.ContactValue,
		OptedIn:      false,
	}

	c, err := d.create.GetOrCreateContact(ctx, ct)
	if err != nil {
		return nil, err
	}

	return &domain.Contact{
		ID:           c.ContactID,
		ContactType:  *c.ContactID,
		ContactValue: c.ContactValue,
		Active:       c.Active,
		OptedIn:      c.OptedIn,
	}, nil
}

// CreateAppointment creates a new appointment
func (d *MyCareHubDb) CreateAppointment(ctx context.Context, appointment domain.Appointment) error {

	date := appointment.Date.AsTime()
	ap := &gorm.Appointment{
		Active:     true,
		ExternalID: appointment.ExternalID,
		ClientID:   appointment.ClientID,
		FacilityID: appointment.FacilityID,
		Reason:     appointment.Reason,
		Provider:   appointment.Provider,
		Date:       date,
	}

	return d.create.CreateAppointment(ctx, ap)
}

// AnswerScreeningToolQuestions creates a screening tool answers
func (d *MyCareHubDb) AnswerScreeningToolQuestions(ctx context.Context, screeningToolResponses []*dto.ScreeningToolQuestionResponseInput) error {

	var screeningToolResponsesObj []*gorm.ScreeningToolsResponse
	for _, st := range screeningToolResponses {
		stq := &gorm.ScreeningToolsResponse{
			ClientID:   st.ClientID,
			QuestionID: st.QuestionID,
			Response:   st.Response,
			Active:     true,
		}
		screeningToolResponsesObj = append(screeningToolResponsesObj, stq)
	}
	err := d.create.AnswerScreeningToolQuestions(ctx, screeningToolResponsesObj)
	if err != nil {
		helpers.ReportErrorToSentry(err)
		return err
	}
	return nil
}
