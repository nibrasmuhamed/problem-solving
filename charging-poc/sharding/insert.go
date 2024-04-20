package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	CustomerID     int       `faker:"-" bson:"customer_id"`
	EventDate      time.Time `faker:"-" bson:"event_date"`
	Quantity       int       `faker:"-" bson:"quantity"`
	Amount         float64   `faker:"-" bson:"amount"`
	ExtractedMonth int       `faker:"-" bson:"extracted_month"`
}

func main() {
	// Set up MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:30000"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Choose database and collection
	db := client.Database("testdb")
	collection := db.Collection("events")

	// Set up faker package
	rand.Seed(time.Now().UnixNano())

	// Generate events for January
	for i := 1; i < 5000; i++ {
		event := Event{
			CustomerID: i,
			EventDate:  faker.DateRange(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)),
			Quantity:   rand.Intn(10),
			Amount:     rand.Float64() * 100,
		}
		// Calculate extracted month from event date
		event.ExtractedMonth = int(event.EventDate.Month())
		_, err := collection.InsertOne(ctx, event)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check inserted documents
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var events []Event
	if err = cursor.All(ctx, &events); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %d events\n", len(events))

	fmt.Println("Done")
}

// sh.addShardTag("shard1rs", "jan-jun")
// sh.addShardTag("shard2rs", "jul-dec")
// sh.shardCollection("testdb.events", { "event_date": 1 }, false, { "jan-jun": { "event_date": { $lt: new Date("2022-07-01") } }, "jul-dec": { "event_date": { $gte: new Date("2022-07-01") } } })

// sh.addTagRange(
// 	"testdb.my",
// 	{ "event_date": ISODate("0000-01-01T00:00:00.000Z") },
// 	{ "event_date": ISODate("0000-06-30T23:59:59.999Z") },
// 	"jan-jun"
//   )
// sh.addTagRange(
// 	"testdb.my",
// 	{ "event_date": ISODate("0000-07-01T00:00:00.000Z") },
// 	{ "event_date": ISODate("0000-12-31T23:59:59.999Z") },
// 	"jul-dec"
//   )
