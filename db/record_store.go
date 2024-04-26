package db

import (
	"context"

	"github.com/ryanzola/week-17/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordStore interface {
	GetRecords(context.Context) ([]*types.Record, error)
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

func (s *MongoRecordStore) GetRecords(ctx context.Context) ([]*types.Record, error) {
	return nil, nil
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
