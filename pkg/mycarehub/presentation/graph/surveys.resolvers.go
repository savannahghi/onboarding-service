package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/dto"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
)

// SendClientSurveyLinks is the resolver for the sendClientSurveyLinks field.
func (r *mutationResolver) SendClientSurveyLinks(ctx context.Context, facilityID string, formID string, projectID int, filterParams *dto.ClientFilterParamsInput) (bool, error) {
	return r.mycarehub.Surveys.SendClientSurveyLinks(ctx, &facilityID, &formID, &projectID, filterParams)
}

// VerifySurveySubmission is the resolver for the verifySurveySubmission field.
func (r *mutationResolver) VerifySurveySubmission(ctx context.Context, input dto.VerifySurveySubmissionInput) (bool, error) {
	return r.mycarehub.Surveys.VerifySurveySubmission(ctx, input)
}

// ListSurveys is the resolver for the listSurveys field.
func (r *queryResolver) ListSurveys(ctx context.Context, projectID int) ([]*domain.SurveyForm, error) {
	return r.mycarehub.Surveys.ListSurveys(ctx, &projectID)
}

// GetUserSurveyForms is the resolver for the getUserSurveyForms field.
func (r *queryResolver) GetUserSurveyForms(ctx context.Context, userID string) ([]*domain.UserSurvey, error) {
	return r.mycarehub.Surveys.GetUserSurveyForms(ctx, userID)
}

// ListSurveyRespondents is the resolver for the listSurveyRespondents field.
func (r *queryResolver) ListSurveyRespondents(ctx context.Context, projectID int, formID string, paginationInput dto.PaginationsInput) (*domain.SurveyRespondentPage, error) {
	return r.mycarehub.Surveys.ListSurveyRespondents(ctx, projectID, formID, paginationInput)
}

// GetSurveyServiceRequestUser is the resolver for the getSurveyServiceRequestUser field.
func (r *queryResolver) GetSurveyServiceRequestUser(ctx context.Context, facilityID string, projectID int, formID string, paginationInput dto.PaginationsInput) (*domain.SurveyServiceRequestUserPage, error) {
	return r.mycarehub.Surveys.GetSurveyServiceRequestUser(ctx, facilityID, projectID, formID, paginationInput)
}

// GetSurveyResponse is the resolver for the getSurveyResponse field.
func (r *queryResolver) GetSurveyResponse(ctx context.Context, input dto.SurveyResponseInput) ([]*domain.SurveyResponse, error) {
	return r.mycarehub.Surveys.GetSurveyResponse(ctx, input)
}

// GetSurveyWithServiceRequest is the resolver for the getSurveyWithServiceRequest field.
func (r *queryResolver) GetSurveyWithServiceRequest(ctx context.Context, facilityID string) ([]*dto.SurveysWithServiceRequest, error) {
	return r.mycarehub.Surveys.GetSurveysWithServiceRequests(ctx, facilityID)
}
