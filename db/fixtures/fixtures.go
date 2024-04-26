package fixtures

import (
	"context"
	"log"
	"time"

	"github.com/ryanzola/week-17/db"
	"github.com/ryanzola/week-17/types"
)

func AddRecord(store db.RecordStore, key string, createdAt time.Time, totalCount int64) *types.Record {
	record := types.Record{
		Key:        key,
		CreatedAt:  createdAt,
		TotalCount: totalCount,
	}

	insertedRecord, err := store.InsertRecord(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}

	return insertedRecord
}
