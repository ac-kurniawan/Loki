package attendee

type AttendeeCreatedSchema struct {
	EventId       string `json:"eventId"`
	ScheduleIndex int    `json:"scheduleIndex"`
}

type Publish interface {
	AttendeeCreated(schema AttendeeCreatedSchema) error
}
