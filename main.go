package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ryanzola/week-17/api"
	"github.com/ryanzola/week-17/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	var (
		recordStore   = db.NewMongoRecordStore(client)
		recordHandler = api.NewRecordHandler(recordStore)
		memoryHandler = api.NewMemoryHandler()
	)

	http.HandleFunc("POST /api/v1/record", api.Make(recordHandler.HandleGetRecords))
	http.HandleFunc("GET /api/v1/in-memory", memoryHandler.HandleGetMemoryRecords)
	http.HandleFunc("POST /api/v1/in-memory", memoryHandler.HandlePostMemoryRecord)

	listen_addr := os.Getenv("LISTEN_ADDR")
	log.Printf("Starting server on %s", listen_addr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listen_addr), nil))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db.DBNAME = os.Getenv("MONGO_DB_NAME")
	db.DBURI = os.Getenv("MONGO_URI")
}
