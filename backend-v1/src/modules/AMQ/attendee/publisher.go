package AMQ

import (
	"antriin/src/business/attendee"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Publish struct {
	publisher *amqp.Publisher
}

func (p Publish) AttendeeCreated(schema attendee.AttendeeCreatedSchema) error {
	payload, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	return p.publisher.Publish("attendeeCreated", msg)
}

func NewAttendeePublisher(publisher *amqp.Publisher) attendee.Publish {
	return &Publish{
		publisher: publisher,
	}
}