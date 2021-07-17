package attendee

// Attendee entity
type Attendee struct {
	ID       string
	Name     string
	EventId  string
	Contacts []Contact
}

type Contact struct {
	ContactName string
	Contact     string
}
