package mongoRepository

import "antriin/src/business/attendee"

type AttendeeMongo struct {
	ID       string         `bson:"_id,omitempty" `
	Name     string         `bson:"name"`
	EventId  string         `bson:"event_id"`
	Contacts []ContactMongo `bson:"contacts"`
}
type ContactMongo struct {
	ContactName string `bson:"contact_name"`
	Contact     string `bson:"contact"`
}

func NewAttendeeCollection(attendee attendee.Attendee) *AttendeeMongo {
	var contactData []ContactMongo
	for _, e := range attendee.Contacts {
		contactData = append(contactData, ContactMongo{e.ContactName, e.Contact})
	}

	return &AttendeeMongo{
		ID:       attendee.ID,
		Name:     attendee.Name,
		EventId:  attendee.EventId,
		Contacts: contactData,
	}
}

func ToAttendee(data AttendeeMongo) attendee.Attendee {
	var contactData []attendee.Contact
	for _, e := range data.Contacts {
		contactData = append(contactData, attendee.Contact{
			ContactName: e.ContactName,
			Contact:     e.Contact})
	}
	return attendee.Attendee{
		ID:       data.ID,
		Name:     data.Name,
		EventId:  data.EventId,
		Contacts: contactData,
	}
}
