package mongoRepository

import (
	"antriin/src/business/event"
	"antriin/src/util/debug"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (r Repository) GetEventsByCreatorID(creatorID string) ([]event.Event,
	error) {
	opts := options.Find().SetLimit(20).SetSkip(0)
	docs, err := r.collection.Find(context.TODO(), bson.D{{"creator_id",
		creatorID}},
		opts)
	if err != nil {
		debug.Error("eventRepository", err.Error())
		return []event.Event{}, err
	}

	// data mapping
	var result []event.Event
	for docs.Next(context.TODO()) {
		var data EventMongo
		err := docs.Decode(&data)
		if err != nil {
			debug.Error("EventRepository", err.Error())
			return nil, err
		}
		result = append(result, ToEvent(data))
	}

	if err := docs.Err(); err != nil {
		debug.Error("EventRepository", err.Error())
		return nil, err
	}

	if err := docs.Close(context.TODO()); err != nil {
		debug.Error("EventRepository", err.Error())
		return nil, err
	}

	return result, nil
}

func (r Repository) SetEvent(data event.Event, creatorID string) (event.Event,
	error) {
	data.CreatorID = creatorID
	col := NewEventCollection(data)

	res, err := r.collection.InsertOne(context.TODO(), col)
	if err != nil {
		debug.Error("EventRepository", err.Error())
		return event.Event{}, err
	}

	data.ID = res.InsertedID.(primitive.ObjectID).String()

	return data, nil
}

func (r Repository) GetEventById(id string) (event.Event, error) {
	opts := options.FindOne()
	var result EventMongo
	objcId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		debug.Error("eventRepository", err.Error())
		return event.Event{}, err
	}

	err = r.collection.FindOne(
		context.TODO(), bson.M{
			"_id": bson.M{"$eq": objcId},
		},
		opts,
	).Decode(&result)

	if err != nil {
		debug.Error("eventRepository", err.Error())
		return event.Event{}, err
	}

	res := ToEvent(result)
	return res, nil
}

func (r Repository) CreateSchedules(
	eventId string,
	schedules []event.Schedule,
) (event.Event, error) {
	objcId, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		debug.Error("eventRepository", err.Error())
		return event.Event{}, err
	}

	updated, err := r.collection.UpdateOne(context.Background(),
		bson.M{"_id": bson.M{"$eq": objcId}}, bson.M{
			"$set": bson.M{
				"schedule": NewScheduleCollection(schedules),
			},
		})

	if err != nil {
		debug.Error("eventRepository", err.Error())
		return event.Event{}, err
	}

	if updated.MatchedCount <= 0 {
		return event.Event{}, errors.New("Event not found")
	}

	return r.GetEventById(eventId)
}

func NewEventRepository(client *mongo.Client) event.IEventRepository {
	collection := client.Database("antri").Collection("events")
	return &Repository{client: client, collection: collection}
}
