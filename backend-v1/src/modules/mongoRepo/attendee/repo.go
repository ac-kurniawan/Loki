package mongoRepository

import (
	"antriin/src/business/attendee"
	"antriin/src/util/debug"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	collection *mongo.Collection
}

func (r Repository) SetAttendee(data attendee.Attendee) error {
	col := NewAttendeeCollection(data)
	_, err := r.collection.InsertOne(context.TODO(), col)
	if err != nil {
		debug.Error("AttendeeRepository", err.Error())
		return err
	}

	return nil
}

func (r Repository) GetAttendees(page int) ([]attendee.Attendee, error) {
	skip := 20 * (page - 1)
	findOpt := options.Find().SetLimit(20).SetSkip(int64(skip))

	docs, err := r.collection.Find(context.TODO(), bson.M{}, findOpt)
	if err != nil {
		debug.Error("AttendeeRepository", err.Error())
		return nil, err
	}

	// data mapping
	var result []attendee.Attendee
	for docs.Next(context.TODO()) {
		var data AttendeeMongo
		err := docs.Decode(&data)
		if err != nil {
			debug.Error("AttendeeRepository", err.Error())
			return nil, errors.New(err.Error())
		}
		result = append(result, ToAttendee(data))
	}

	if err := docs.Err(); err != nil {
		debug.Error("AttendeeRepository", err.Error())
		return nil, errors.New(err.Error())
	}

	if err := docs.Close(context.TODO()); err != nil {
		debug.Error("AttendeeRepository", err.Error())
		return nil, errors.New(err.Error())
	}

	return result,nil
}

func NewAttendeeRepository(client *mongo.Client) attendee.IAttendeeRepository {
	collection := client.Database("antri").Collection("attendees")

	return &Repository{client, collection}
}