package model

// RecordID defines a record id. Together with RecordType identifies unique records across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID identifies unique records across all types.
type RecordType string

// Existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int

// Rating defines an individual rating created by a user for some record.
type Rating struct {
	RecordID   string      `json:"recordId" validate:"required"`
	RecordType string      `json:"recordType" validate:"required"`
	UserID     UserID      `json:"userId" validate:"required"`
	Value      RatingValue `json:"value" validate:"gte=1,lte=5"`
}

// RatingEvent defines an event containing rating information.
type RatingEvent struct {
	UserID     UserID          `json:"userId" validate:"required"`
	RecordID   RecordID        `json:"recordId" validate:"required"`
	RecordType RecordType      `json:"recordType" validate:"required"`
	Value      RatingValue     `json:"value" validate:"gte=1,lte=5"`
	ProviderID string          `json:"providerId" validate:"required"`
	EventType  RatingEventType `json:"eventType" validate:"required"`
}

// RatingEventType defines the type of a rating event.
type RatingEventType string

// Rating event types.
const (
	RatingEventTypePut    = RatingEventType("put")
	RatingEventTypeDelete = RatingEventType("delete")
)
