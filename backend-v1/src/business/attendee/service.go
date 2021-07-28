package attendee

import (
	"errors"
)


type IAttendee interface {
	SetAttendee(data Attendee) error
	GetAttendees(page int) ([]Attendee,error)
	GetAttendeeById(id int) (*Attendee,error)
}

// Service AttendeeService constructor and method
type Service struct {
	attendeeRepository IAttendeeRepository
	publisher Publish
}

func (a Service) GetAttendeeById(id int) (*Attendee, error) {
	return nil,errors.New("no method")
}

func (a Service) SetAttendee(data Attendee) error {
	err := a.attendeeRepository.SetAttendee(data)
	err = a.publisher.AttendeeCreated(AttendeeCreatedSchema{
		EventId: data.EventId,
		ScheduleIndex: 0,
	})
	return err
}

func (a Service) GetAttendees(page int) ([]Attendee, error) {
	col, err := a.attendeeRepository.GetAttendees(page)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// NewAttendee attendee builder service
func NewAttendee(attendeeRepository IAttendeeRepository,
	publisher Publish) IAttendee  {
	return &Service{attendeeRepository, publisher}
}