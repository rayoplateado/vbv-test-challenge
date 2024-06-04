package domain

import (
	"factorialapi/database"
	"factorialapi/log"
	"time"
)

var databaseName = "factorial_codechallenge"
var collectionName = "metrics"

type Metric struct {
	Id         string                       `json:"_id" bson:"_id,omitempty"`
	Timestamp  time.Time                    `json:"Timestamp" bson:"Timestamp,omitempty"`
	Name       string                       `json:"Name" bson:"Name,omitempty"`
	Value      string                       `json:"Value" bson:"Value,omitempty"`
	Collection database.ICollection[Metric] `bson:"-"`
}

func (metric Metric) List() ([]Metric, error) {
	log.Info("listing data")
	metrics, err := metric.Collection.List(databaseName, collectionName)

	if err != nil {
		log.Error("error listing data: " + err.Error())
		return nil, err
	}

	return metrics, nil
}

func (metric Metric) Find(id string) (Metric, error) {
	log.Info("listing data by id")
	metric, err := metric.Collection.FindOne(databaseName, collectionName, id)

	if err != nil {
		log.Error("error listing data: " + err.Error())
		return Metric{}, err
	}

	return metric, nil
}

func (metric Metric) FindList() ([]Metric, error) {
	log.Info("listing data by filter")
	metrics, err := metric.Collection.Find(databaseName, collectionName, metric)

	if err != nil {
		log.Error("error listing data: " + err.Error())
		return nil, err
	}

	return metrics, nil
}

func (metric Metric) Add() (Metric, error) {
	log.Info("adding data")
	id, err := metric.Collection.Add(databaseName, collectionName, metric)

	if err != nil {
		log.Error("error adding data: " + err.Error())
		return Metric{}, err
	}

	metric.Id = id

	return metric, nil
}
