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
		key := generateKey(16)
		createdAt := generateRandomDate(10)
		totalCount := int64(i) + 1

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

func generateKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length) // generate len bytes of random data
	rand.Read(bytes)

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))] // map each byte to a character in the charset
	}

	return string(bytes)
}

func generateRandomDate(yearsAgo int) time.Time {
	now := time.Now()
	tenYearsAgo := now.AddDate(-yearsAgo, 0, 0)
	days := now.Sub(tenYearsAgo).Hours() / 24

	randomDays := rand.Intn(int(days))
	return now.AddDate(0, 0, -randomDays)
}
