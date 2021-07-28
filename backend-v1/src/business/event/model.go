package event

import "time"

// Event Entity
type Event struct {
	ID        string
	Name      string
	CreatorID string
	Schedule  []Schedule
}

type Schedule struct {
	Location     Location
	Date         time.Time
	Start        time.Time
	End          time.Time
	Capacity     uint
	Progress     uint
	AttendeeType string
}

type Location struct {
	Address     string
	District    string
	SubDistrict string
	City        string
	Province    string
	Longitude   float64
	Latitude    float64
}
