package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Choose your database and collection
	database := client.Database("balance")
	collection := database.Collection("Customers")

	// Define your query (you can customize this based on your requirements)
	query := bson.D{{Key: "customerid", Value: "cust187501"}, {Key: "operatorid", Value: "e51da9b9-259e-421c-8bf6-22f5bdd20406"}}

	// Create an explain command
	explainCommand := bson.D{
		{"explain", bson.M{
			"find":   "Customers",
			"filter": query,
		}},
	}

	// Execute the explain command
	var result bson.M
	err = collection.Database().RunCommand(ctx, explainCommand).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	prettyJson, err := json.MarshalIndent(result, "", "  ")
	// Print the explain result
	fmt.Printf("Explain Plan: %+v\n", string(prettyJson))

	fmt.Printf("%v\n", result["executionStats"])
	executionStatsVal, found := result["executionStats"].(primitive.M)
	if !found {
		fmt.Println("Unable to determine if an index is used.")
		return
	}

	usedIndex, ok := executionStatsVal["totalKeysExamined"].(int32)
	if ok {
		fmt.Println("Total Keys scanned is ", usedIndex)
	}
}
