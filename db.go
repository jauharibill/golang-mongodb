package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var red *redis.Client
var mongodb *mongo.Client
var err error

func init() {
	mongodb, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	red = redis.NewClient(&redis.Options{DB: 0, Password: "", Addr: ":6379"})
}

func Conn() *mongo.Client {
	err = mongodb.Connect(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}

	return mongodb
}
