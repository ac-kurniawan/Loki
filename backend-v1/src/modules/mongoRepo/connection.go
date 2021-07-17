package mongoRepo

import (
	"antriin/src/util/debug"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connect() (*mongo.Client, error){
	credential := options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	opt := options.Client().ApplyURI(fmt.Sprintf("%s:%s",
		os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))).SetAuth(
		credential)

	client, err := mongo.Connect(context.Background(), opt)

	if err != nil {
		debug.Error("MongoConnection", err.Error())
		return nil, errors.New(err.Error())
	}

	debug.Info("MongoConnection", "connection established")
	return client,nil
}