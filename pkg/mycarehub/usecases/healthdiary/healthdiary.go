package healthdiary

import (
	"context"
	"fmt"
	"time"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/domain"
	"github.com/savannahghi/mycarehub/pkg/mycarehub/infrastructure"
)

// The healthdiary is used for engagement with clients on a day-by-day basis.
// The idea behind it is to track sustained changes in client's moods. The moods range
// from VERY_HAPPY, HAPPY, NEUTRAL, SAD, VERY_SAD. When a client fills the mood board, a health diary
// entry is recorded in the database. In cases where the client is VERY_SAD, the client is asked if they
// want to report it to a healthcare worker and if they do, a service request is created. The service request
// is a task for the healthcare worker on the platform. All this should happen within a 24 hour time window. If
// a health diary was filled within the past 24 hours, the client is shown an inspirational post on the frontend
// and if it hasn't been filled, we show them the health diary.

// CreateHealthDiaryEntry is an interface that holds the method signature for creating a health diary entry
type CreateHealthDiaryEntry interface {
	CreateHealthDiaryEntry(ctx context.Context, clientID string, note *string, mood string, reportToStaff bool) (bool, error)
}

// UseCasesHealthDiary holds all the interfaces that represents the business logic to implement the health diary
type UseCasesHealthDiary interface {
	CreateHealthDiaryEntry
}

// UseCasesHealthDiaryImpl embeds the healthdiary logic defined on the domain
type UseCasesHealthDiaryImpl struct {
	Create infrastructure.Create
}

// NewUseCaseHealthDiaryImpl creates a new instance of health diary
func NewUseCaseHealthDiaryImpl(
	create infrastructure.Create,
) *UseCasesHealthDiaryImpl {
	return &UseCasesHealthDiaryImpl{
		Create: create,
	}
}

// CreateHealthDiaryEntry captures a client's mood and creates a health diary entry. This will be used to
// track the client's moods on a day-to-day basis
func (h UseCasesHealthDiaryImpl) CreateHealthDiaryEntry(
	ctx context.Context,
	clientID string,
	note *string,
	mood string,
	reportToStaff bool,
) (bool, error) {
	switch mood {
	case string(enums.MoodVerySad):
		currentTime := time.Now()
		healthDiaryEntry := &domain.ClientHealthDiaryEntry{
			Active:                true,
			Mood:                  mood,
			Note:                  *note,
			EntryType:             "HOME_PAGE_HEALTH_DIARY_ENTRY", //TODO: Make this an enum
			ShareWithHealthWorker: reportToStaff,
			ClientID:              clientID,
			SharedAt:              &currentTime,
		}

		serviceRequest := &domain.ClientServiceRequest{
			Active:       true,
			RequestType:  "HEALTH_DIARY_ENTRY", //TODO make this an enum
			Request:      "",
			Status:       "PENDING", // TODO; enum
			InProgressAt: time.Now(),
			ClientID:     clientID,
		}

		err := h.Create.CreateServiceRequest(ctx, healthDiaryEntry, serviceRequest)
		if err != nil {
			return false, fmt.Errorf("failed to create service request: %v", err)
		}

	default:
		healthDiaryEntry := &domain.ClientHealthDiaryEntry{
			Active:                true,
			Mood:                  mood,
			Note:                  *note,
			EntryType:             "HOME_PAGE_HEALTH_DIARY_ENTRY", //TODO: Make this an enum
			ShareWithHealthWorker: false,
			ClientID:              clientID,
			SharedAt:              nil,
		}
		err := h.Create.CreateHealthDiaryEntry(ctx, healthDiaryEntry)
		if err != nil {
			return false, fmt.Errorf("failed to save health diary entry")
		}
	}
	return true, nil
}
