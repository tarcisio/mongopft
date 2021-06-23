package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDocument struct {
	ID            string `bson:"id,omitempty"`
	Title         string `bson:"title,omitempty"`
	Body          string `bson:"body,omitempty"`
	UnixTimestamp int64  `bson:"timestamp,omitempty"`
}

func NewTestDocument() TestDocument {

	return TestDocument{
		ID:            RandStringBytes(50),
		Title:         RandStringBytes(RandNumber()),
		Body:          RandStringBytes(500),
		UnixTimestamp: time.Now().Unix(),
	}
}

func TestThread(ctx context.Context, dsn string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("db01").Collection("collection01")

	mychar := randCliChar()
	for {
		select {
		case <-time.After(1 * time.Millisecond):
			dc := NewTestDocument()
			_, err := collection.InsertOne(ctx, dc)
			if err != nil {
				return err
			}
			fmt.Println(mychar)
		case <-ctx.Done():
			fmt.Println("halted operation")
		}
	}
	return nil
}
