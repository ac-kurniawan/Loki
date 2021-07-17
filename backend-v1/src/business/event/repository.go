package event

type IEventRepository interface {
	SetEvent(data Event, creatorId string) (Event, error)
	GetEventsByCreatorID(creatorID string) ([]Event, error)
	GetEventById(id string) (Event,error)

	CreateSchedules(eventId string, schedules []Schedule) (Event,error)
}