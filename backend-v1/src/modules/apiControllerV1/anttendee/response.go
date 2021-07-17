package apiControllerV1

import "antriin/src/business/attendee"

type GetAttendeeResponse struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	EventId  string            `json:"eventId,omitempty"`
	Contacts []ContactResponse `json:"contacts"`
}
type ContactResponse struct {
	ContactName string `json:"contactName"`
	Contact     string `json:"contact"`
}

func NewGetAttendeesResponse(data []attendee.Attendee) []GetAttendeeResponse {
	var buffer []GetAttendeeResponse

	for _, e := range data {

		var contactData []ContactResponse
		for _, e2 := range e.Contacts {
			contactData = append(contactData, ContactResponse{
				e2.ContactName,
				e2.Contact,
			})
		}

		buffer = append(buffer, GetAttendeeResponse{
			e.ID,
			e.Name,
			e.EventId,
			contactData,
		})
	}

	return buffer
}
