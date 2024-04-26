package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Key        string             `bson:"key" json:"key"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	TotalCount int64              `bson:"totalCount" json:"totalCount"`
}
