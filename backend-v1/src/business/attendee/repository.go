package attendee

type IAttendeeRepository interface {
	GetAttendees(page int) ([]Attendee,error)
	SetAttendee(data Attendee) error
}
