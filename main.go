package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ryanzola/week-17/api"
	"github.com/ryanzola/week-17/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		app           = fiber.New(config)
		apiv1         = app.Group("/api/v1")
		recordStore   = db.NewMongoRecordStore(client)
		recordHandler = api.NewRecordHandler(recordStore)
	)

	// GET /api/v1/hello
	apiv1.Get("/record", recordHandler.HandleGetRecords)
	apiv1.Get("/record/:id", recordHandler.HandleGetRecordByID)
	apiv1.Post("/record", recordHandler.HandleInsertRecord)
	apiv1.Put("/record/:id", recordHandler.HandleUpdateRecord)
	apiv1.Delete("/record/:id", recordHandler.HandleDeleteRecord)

	listen_addr := os.Getenv("LISTEN_ADDR")
	log.Printf("Starting server on %s", listen_addr)
	app.Listen(listen_addr)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db.DBNAME = os.Getenv("MONGO_DB_NAME")
	db.DBURI = os.Getenv("MONGO_URI")
}
