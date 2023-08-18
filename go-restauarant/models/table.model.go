package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct{
	ID				primitive.ObjectID			`bson:"_id"`
	Updated_at		time.Time					`json:"updated_id"`
	Table_id		string						`json:"table_id"`
}