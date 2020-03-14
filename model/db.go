package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client
var ctx context.Context
var cancel func()

func init() {
	var err error

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	db, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalln(err)
	}

	// Ping mongo database if up
	go func() {
		for {
			if err := db.Ping(ctx, readpref.Primary()); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 5)
		}
	}()
}
