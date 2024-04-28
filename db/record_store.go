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
	GetRecords(context.Context, types.RecordQueryParams) ([]*types.Record, error)
	GetRecordByID(context.Context, string) (*types.Record, error)
	InsertRecord(context.Context, *types.Record) (*types.Record, error)
	UpdateRecord(context.Context, *types.Record) error
	DeleteRecord(context.Context, string) error
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

func (s *MongoRecordStore) GetRecords(ctx context.Context, filter types.RecordQueryParams) ([]*types.Record, error) {
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

func (s *MongoRecordStore) GetRecordByID(ctx context.Context, id string) (*types.Record, error) {
	return nil, nil
}

func (s *MongoRecordStore) InsertRecord(ctx context.Context, record *types.Record) (*types.Record, error) {
	resp, err := s.coll.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	record.ID = resp.InsertedID.(primitive.ObjectID)
	return record, nil
}

func (s *MongoRecordStore) UpdateRecord(ctx context.Context, record *types.Record) error {
	return nil
}

func (s *MongoRecordStore) DeleteRecord(ctx context.Context, id string) error {
	return nil
}
