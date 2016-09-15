package queue

import "time"

type Message struct {
	ID         int64                  `json:"id" bson:"_id"`      // The Messages ID
	Queue      string                 `json:"queue" bson:"queue"` // The Queue ID for the message
	Data       map[string]interface{} `json:"data" bson:"data"`   // The data for the Message
	Date       time.Time              `json:"-" bson:"date"`      // The date the message was queued
	InProgress bool                   `json:"-" bson:"_inProg"`   // Indicator if this messages is in progress
	ExpireDate time.Time              `json:"-" bson:"_exp"`      // When this message needs to be done by
	Done       bool                   `json:"-" bson:"_done"`     // Whether or not this message is complete
}
