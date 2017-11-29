package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoDBDatastore struct {
	*mgo.Session
}

func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoDBDatastore{
		Session: session,
	}, nil
}

func (m *MongoDBDatastore) createActivity(a Activity) error {
	session := m.Copy()

	defer session.Close()

	activityCollection := session.DB("timeline").C("activities")
	err := activityCollection.Insert(a)

	return err

}

func (m *MongoDBDatastore) completeActivity(id string) error {
	session := m.Copy()
	defer session.Close()

	now := time.Now()
	unixTimeNow := now.Unix()
	activityCollection := session.DB("timeline").C("activities")
	err := activityCollection.Update(bson.M{"_id": id}, bson.M{"timestampCompleted": unixTimeNow})

	return err
}

func (m *MongoDBDatastore) getUserTimeLine(user User) (*[]Activity, error) {
	session := m.Copy()

	defer session.Close()
	activityCollection := session.DB("timeline").C("activities")

	activities := []Activity{}
	err := activityCollection.Find(bson.M{"firstName": user.FirstName, "lastName": user.LastName}).All(&activities)
	if err != nil {
		return nil, err
	}

	return &activities, nil
}

func (m *MongoDBDatastore) Close() {
	m.Close()
}
