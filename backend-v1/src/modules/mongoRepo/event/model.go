package mongoRepository

import (
	"antriin/src/business/event"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EventMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	CreatorID string             `bson:"creator_id"`
	Schedule  []ScheduleMongo    `bson:"schedule"`
}

type ScheduleMongo struct {
	Location     LocationMongo `bson:"location"`
	Date         time.Time     `bson:"date"`
	Start        time.Time     `bson:"start"`
	End          time.Time     `bson:"end"`
	Capacity     uint          `bson:"capacity"`
	Progress     uint          `bson:"progress"`
	AttendeeType string        `bson:"attendee_type"`
}

type LocationMongo struct {
	Address     string  `bson:"address"`
	District    string  `bson:"district"`
	SubDistrict string  `bson:"sub_district"`
	City        string  `bson:"city"`
	Province    string  `bson:"province"`
	Longitude   float64 `bson:"longitude"`
	Latitude    float64 `bson:"latitude"`
}

func ToEvent(data EventMongo) event.Event {
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
			Progress:     e.Progress,
			AttendeeType: e.AttendeeType,
		})
	}
	return event.Event{
		ID:        data.ID.Hex(),
		Name:      data.Name,
		CreatorID: data.CreatorID,
		Schedule:  scheduleData,
	}
}

func NewScheduleCollection(data []event.Schedule) []ScheduleMongo {
	var scheduleData []ScheduleMongo
	for _, e := range data {
		scheduleData = append(scheduleData, ScheduleMongo{
			Location: LocationMongo{
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
			Progress:     e.Progress,
			AttendeeType: e.AttendeeType,
		})
	}
	return scheduleData
}

func NewEventCollection(data event.Event) EventMongo {
	scheduleData := NewScheduleCollection(data.Schedule)
	id, _ := primitive.ObjectIDFromHex(data.ID)
	return EventMongo{
		ID:        id,
		Name:      data.Name,
		CreatorID: data.CreatorID,
		Schedule:  scheduleData,
	}
}
