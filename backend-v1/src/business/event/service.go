package event

type IEvent interface {
	SetEvent(data Event, creatorID string) (Event, error)
	GetEventsByCreatorID(creatorID string) (int32, []Event, error)
	GetEventById(id string) (Event,error)

	CreateSchedules(eventId string, schedules []Schedule) (Event,error)

	AddAttendeeInSchedule(eventId string, scheduleIndex int) error

}

// Service CommitteeService constructor and method
type Service struct {
	repo IEventRepository
}

func (s Service) AddAttendeeInSchedule(eventId string,
	scheduleIndex int) error {
	events, err := s.repo.GetEventById(eventId)
	if err != nil {
		return err
	}

	events.Schedule[scheduleIndex].Progress += 1

	_, err = s.repo.CreateSchedules(eventId, events.Schedule)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetEventsByCreatorID(creatorID string) (int32, []Event,
	error) {
	return s.repo.GetEventsByCreatorID(creatorID)
}

func (s Service) CreateSchedules(eventId string, schedules []Schedule) (Event,
	error) {
	return s.repo.CreateSchedules(eventId, schedules)
}

func (s Service) SetEvent(data Event, creatorID string) (Event, error) {
	return s.repo.SetEvent(data, creatorID)
}

func (s Service) GetEventById(id string) (Event, error) {
	return s.repo.GetEventById(id)
}

// NewEvent event builder service
func NewEvent(repo IEventRepository) IEvent  {
	return &Service{repo}
}
