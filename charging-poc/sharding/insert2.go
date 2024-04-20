package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	CustomerID   int       `bson:"customer_id"`
	CreatedDate  time.Time `bson:"created_date"`
	Quantity     int       `bson:"quantity"`
	Amount       float64   `bson:"amount"`
	ExtractedMon int       `bson:"extracted_month"`
}

func main() {
	// Set up MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:30000"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("testdb").Collection("events")

	// Set up random seed
	rand.Seed(time.Now().UnixNano())

	// Generate random events for different times in a year
	for i := 0; i < 1000; i++ {
		// Generate random customer ID
		customerID := gofakeit.Number(1, 100)

		// Generate random created date
		year := strconv.Itoa(time.Now().Year())
		month := strconv.Itoa(gofakeit.Number(1, 12))
		day := strconv.Itoa(gofakeit.Number(1, 28))
		dateStr := fmt.Sprintf("%s-%s-%s", year, month, day)
		createdDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
		}

		// Generate random quantity and amount
		quantity := gofakeit.Number(1, 10)
		amount := gofakeit.Float64Range(10, 1000)

		// Extract month from created date
		extractedMon := createdDate.Month()

		// Create event object
		event := Event{
			CustomerID:   customerID,
			CreatedDate:  createdDate,
			Quantity:     quantity,
			Amount:       amount,
			ExtractedMon: int(extractedMon),
		}

		// Insert event into MongoDB
		_, err = collection.InsertOne(ctx, event)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Events generated and stored in MongoDB.")
}
