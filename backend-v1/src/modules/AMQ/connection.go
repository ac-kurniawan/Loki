package AMQ

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)


type Config struct {
	AmqpConfig amqp.Config
}

func NewAMQP(strConnection string) (*amqp.Publisher, *amqp.Subscriber, error) {
	amqpConfig := amqp.NewDurableQueueConfig(strConnection)
	subscriber, err := amqp.NewSubscriber(
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}
	return publisher, subscriber, nil
}