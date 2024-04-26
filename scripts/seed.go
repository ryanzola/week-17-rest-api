package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ryanzola/week-17/db"
	"github.com/ryanzola/week-17/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var ctx = context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	recordStore := db.NewMongoRecordStore(client)

	for i := 0; i < 200; i++ {
		key := fmt.Sprintf("key-%d", i)
		createdAt := time.Now().Add(time.Duration(rand.Intn(100)) * time.Hour)
		totalCount := rand.Int63n(1000) + 1000

		inserted := fixtures.AddRecord(recordStore, key, createdAt, totalCount)
		fmt.Printf("Inserted record: %v\n", inserted)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db.DBNAME = os.Getenv("MONGO_DB_NAME")
	db.DBURI = os.Getenv("MONGO_URI")
}
