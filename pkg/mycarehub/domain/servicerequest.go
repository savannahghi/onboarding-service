package domain

import (
	"time"

	"github.com/savannahghi/mycarehub/pkg/mycarehub/application/enums"
)

// ServiceRequest is a domain entity that represents a service request.
type ServiceRequest struct {
	ID            string     `json:"id"`
	RequestType   string     `json:"requestType"`
	Request       string     `json:"request"`
	Status        string     `json:"status"`
	ClientID      string     `json:"clientID"`
	CreatedAt     time.Time  `json:"created"`
	InProgressAt  *time.Time `json:"inProgressAt"`
	InProgressBy  *string    `json:"inProgressBy"`
	ResolvedAt    *time.Time `json:"resolvedAt"`
	ResolvedBy    *string    `json:"resolvedBy"`
	FacilityID    *string    `json:"facility_id"`
	ClientName    *string    `json:"client_name"`
	ClientContact *string    `json:"client_contact"`
}

// RequestTypeCount ...
type RequestTypeCount struct {
	RequestType enums.ServiceRequestType `json:"requestType"`
	Total       int                      `json:"total"`
}

// ServiceRequestsCount ...
type ServiceRequestsCount struct {
	Total             int                 `json:"total"`
	RequestsTypeCount []*RequestTypeCount `json:"requestsTypeCount"`
}
