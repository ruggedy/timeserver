package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Activity struct {
	ID                 bson.ObjectId `json:"id" bson:"_id"`
	FirstName          string        `json:"firstName" bson:"firstName"`
	ActivityName       string        `json:"activityName" bson:"activityName"`
	LastName           string        `json:"lastName" bson:"lastName"`
	TimestampCreated   int64         `json:"timestampCreated" bson:"timestampCreated"`
	TimestampCompleted int64         `json:"timestampCompleted" bson:"timestampCompleted"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func newActivity(firstName string, activityName string, lastName string) *Activity {
	now := time.Now()
	unixTimestamp := now.Unix()

	a := Activity{
		FirstName:        firstName,
		LastName:         lastName,
		ActivityName:     activityName,
		TimestampCreated: unixTimestamp,
	}

	return &a
}
