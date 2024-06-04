package domain

import (
	"factorialfrontend/api"
	"time"
)

type Metric struct {
	Id        string    `json:"_id"`
	Timestamp time.Time `json:"Timestamp"`
	Name      string    `json:"Name"`
	Value     string    `json:"Value"`
	API       api.IApi[Metric]
}

func (metric Metric) List() ([]Metric, error) {
	list, err := metric.API.List("/metric")
	return list, err
}

func (metric Metric) FindOne(id string) (Metric, error) {
	metric, err := metric.API.Get("/metric/" + id)
	return metric, err
}

func (metric Metric) Add() (Metric, error) {
	metric, err := metric.API.Post("/metric", metric)
	return metric, err
}
