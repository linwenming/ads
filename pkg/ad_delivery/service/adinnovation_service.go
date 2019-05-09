package service

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var GAdInnovationService *AdInnovationService

type AdInnovation struct {
	MongoId    bson.ObjectId `bson:"_id"`
	UserId     string        `bson:"user_id"`
	Name       string        `bson:"name"`
	Type       int           `bson:"type"`
	Material   int           `bson:"material"`
	Height     int64         `bson:"height"`
	Width      int64         `bson:"width"`
	Size       int           `bson:"size"`
	Duration   int           `bson:"duration"`
	Url        string        `bson:"url"`
	CreateTime int64         `bson:"create_time"`
	UpdateTime int64         `bson:"update_time"`
}

type AdInnovationService struct {
	Collection *mgo.Collection
}

func (s *AdInnovationService) CreateAdInnovation(inno *AdInnovation) (id string, err error) {
	inno.MongoId = bson.NewObjectId()
	inno.CreateTime = time.Now().Unix()
	inno.UpdateTime = time.Now().Unix()
	err = s.Collection.Insert(inno)
	return inno.MongoId.Hex(), err
}