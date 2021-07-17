package apiControllerV1

type SetAttendeeRequest struct {
	Name     string              `json:"name" validate:"required"`
	EventId  string              `json:"eventId" validate:"required"`
	Contacts []SetContactRequest `json:"contacts"`
}

type SetContactRequest struct {
	ContactName string `json:"contactName"`
	Contact     string `json:"contact"`
}
