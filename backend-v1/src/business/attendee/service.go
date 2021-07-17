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
}

func (a Service) GetAttendeeById(id int) (*Attendee, error) {
	return nil,errors.New("no method")
}

func (a Service) SetAttendee(data Attendee) error {
	err := a.attendeeRepository.SetAttendee(data)

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
func NewAttendee(attendeeRepository IAttendeeRepository) IAttendee  {
	return &Service{attendeeRepository}
}