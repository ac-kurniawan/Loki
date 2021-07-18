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

func (r Repository) GetEventsByCreatorID(creatorID string) (int32,
	[]event.Event,
	error) {
	matchStage := bson.D{{"$match", bson.D{{"creator_id", creatorID}}}}
	facedStage := bson.D{{"$facet", bson.D{
		{"count", bson.A{bson.D{{"$count", "value"}}}},
		{"data", bson.A{
			bson.D{{"$sort", bson.D{{"_id", -1}}}},
			bson.D{{"$skip", 0}},
			bson.D{{"$limit", 20}},
		}},
	}}}
	unwindStage := bson.D{{"$unwind", "$count"}}
	setStage := bson.D{{"$set", bson.D{{"count", "$count.value"}}}}

	aggr, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{
		matchStage,
		facedStage,
		unwindStage,
		setStage,
	})
	if err != nil {
		panic(err)
	}

	type dataType struct {
		Count int32         `bson:"count" json:"count"`
		Data  []EventMongo `bson:"data" json:"data"`
	}

	var result []dataType
	if err = aggr.All(context.TODO(), &result); err != nil {
		debug.Error("EventRepository", err.Error())
		return -1, []event.Event{}, err
	}

	var dataEvent []event.Event
	for _, e := range result[0].Data {
		dataEvent = append(dataEvent, ToEvent(e))
	}

	return result[0].Count, dataEvent, nil
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

	data.ID = res.InsertedID.(primitive.ObjectID).Hex()

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
