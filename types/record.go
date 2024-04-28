package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Key        string             `bson:"key" json:"key"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	TotalCount int64              `bson:"totalCount" json:"totalCount"`
}

type ResourceResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Records []*Record `json:"records"`
}

type RecordQueryParams struct {
	StartDate string
	EndDate   string
	MinCount  int64
	MaxCount  int64
}
