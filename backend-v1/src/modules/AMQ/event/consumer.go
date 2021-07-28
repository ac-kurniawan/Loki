package AMQ

import (
	"antriin/src/business/event"
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
)

type EventConsumer interface {
	OnAttendeeCreated()
}

type Handler struct {
	service    event.IEvent
	subscriber *amqp.Subscriber
}
type AttendeeCreated struct {
	EventId       string `json:"eventId"`
	ScheduleIndex int    `json:"scheduleIndex"`
}

func (h Handler) OnAttendeeCreated() {
	logger := watermill.NewStdLogger(false, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddNoPublisherHandler(
		"update-event-on-attendee-created",
		"attendeeCreated",
		h.subscriber,
		func(msg *message.Message) error {
			attendee := AttendeeCreated{}
			err := json.Unmarshal(msg.Payload, &attendee)
			if err != nil {
				return err
			}

			err = h.service.AddAttendeeInSchedule(
				attendee.EventId,
				attendee.ScheduleIndex,
			)
			if err != nil {
				return err
			}

			log.Printf("%s - %+v", string(msg.UUID), attendee)
			return nil
		},
	)

	go func() {
		err := router.Run(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	<-router.Running()
}

func NewEventConsumer(service event.IEvent, subs *amqp.Subscriber) EventConsumer {
	return &Handler{
		service:    service,
		subscriber: subs,
	}
}
