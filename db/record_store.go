package db

import (
	"context"
	"time"

	"github.com/ryanzola/week-17/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Map map[string]any

type RecordStore interface {
	GetRecords(context.Context, types.FilterParamsRequest) ([]*types.Record, error)
	InsertRecord(context.Context, *types.Record) (*types.Record, error)
}

type MongoRecordStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoRecordStore(client *mongo.Client) *MongoRecordStore {
	coll := client.Database(DBNAME).Collection("records")
	return &MongoRecordStore{
		client: client,
		coll:   coll,
	}
}

func (s *MongoRecordStore) GetRecords(ctx context.Context, filter types.FilterParamsRequest) ([]*types.Record, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, filter.StartDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(layout, filter.EndDate)
	if err != nil {
		return nil, err
	}

	where := bson.M{
		"createdAt": bson.M{
			"$gte": start,
			"$lte": end,
		},
		"totalCount": bson.M{
			"$gte": filter.MinCount,
			"$lte": filter.MaxCount,
		},
	}
	resp, err := s.coll.Find(ctx, where)
	if err != nil {
		return nil, err
	}

	var records []*types.Record
	if err := resp.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (s *MongoRecordStore) InsertRecord(ctx context.Context, record *types.Record) (*types.Record, error) {
	resp, err := s.coll.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	record.ID = resp.InsertedID.(primitive.ObjectID)
	return record, nil
}
