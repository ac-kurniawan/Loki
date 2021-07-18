package apiControllerV1

import (
	"antriin/src/business/event"
	"time"
)

type EventResponse struct {
	ID        string             `json:"id,omitempty"`
	Name      string             `json:"name"`
	CreatorID string             `json:"creatorId"`
	Schedule  []ScheduleResponse `json:"schedule"`
}

type ScheduleResponse struct {
	Location     LocationResponse `json:"location"`
	Date         time.Time        `json:"date"`
	Start        time.Time        `json:"start"`
	End          time.Time        `json:"end"`
	Capacity     uint             `json:"capacity"`
	AttendeeType string           `json:"attendeeType"`
}

type LocationResponse struct {
	Address     string  `json:"address"`
	District    string  `json:"district"`
	SubDistrict string  `json:"subDistrict"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

func ResponseToEvent(data EventResponse) event.Event {
	var scheduleData []event.Schedule
	for _, e := range data.Schedule {
		scheduleData = append(scheduleData, event.Schedule{
			Location: event.Location{
				Address:     e.Location.Address,
				District:    e.Location.District,
				SubDistrict: e.Location.SubDistrict,
				City:        e.Location.City,
				Province:    e.Location.Province,
				Longitude:   e.Location.Longitude,
				Latitude:    e.Location.Latitude,
			},
			Date:         e.Date,
			Start:        e.Start,
			End:          e.End,
			Capacity:     e.Capacity,
			AttendeeType: e.AttendeeType,
		})
	}
	return event.Event{
		ID:        data.ID,
		Name:      data.Name,
		CreatorID: data.CreatorID,
		Schedule:  scheduleData,
	}
}

func NewEventResponse(data event.Event) EventResponse {
	var scheduleData []ScheduleResponse
	for _, e := range data.Schedule {
		scheduleData = append(scheduleData, ScheduleResponse{
			Location: LocationResponse{
				Address:     e.Location.Address,
				District:    e.Location.District,
				SubDistrict: e.Location.SubDistrict,
				City:        e.Location.City,
				Province:    e.Location.Province,
				Longitude:   e.Location.Longitude,
				Latitude:    e.Location.Latitude,
			},
			Date:         e.Date,
			Start:        e.Start,
			End:          e.End,
			Capacity:     e.Capacity,
			AttendeeType: e.AttendeeType,
		})
	}
	return EventResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatorID: data.CreatorID,
		Schedule:  scheduleData,
	}
}
