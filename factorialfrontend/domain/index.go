package domain

import (
	"factorialfrontend/api"
	"factorialfrontend/log"
	"time"
)

type Index struct {
	API            api.IApi[Metric]
	Metrics        []Metric
	ChartData      Chart
	IsChartEnabled bool
}

func NewIndex(api api.IApi[Metric]) (Index, error) {
	var index Index
	var metric Metric
	var metrics []Metric

	log.Info("querying list of metrics")
	metric.API = api
	metrics, err := metric.List()

	if err != nil {
		log.Error("error fetching list of data: " + err.Error())
		return Index{}, err
	}

	index.Metrics = metrics
	return index, nil
}

func NewIndexFiltered(api api.IApi[Metric], filterTime string) (Index, error) {
	var index Index
	var metric Metric
	var metrics []Metric
	var filteredList []Metric

	log.Info("querying list of metrics filtered")
	metric.API = api
	metrics, err := metric.List()
	if err != nil {
		log.Error("error fetching list of data: " + err.Error())
		return Index{}, err
	}

	start := time.Now()
	switch filterTime {
	case "minute":
		start = start.Add(-1 * time.Hour)
	case "hour":
		start = start.Add(-1 * time.Hour * 24)
	case "day":
		start = start.Add(-1 * time.Hour * 24 * 30)
	default:
		start = time.Time{}
	}

	for _, item := range metrics {
		if item.Timestamp.After(start) {
			filteredList = append(filteredList, item)
		}
	}

	index.IsChartEnabled = false
	if filterTime != "" {
		chartData := NewChart(filterTime, metrics)
		index.ChartData = chartData
		index.IsChartEnabled = true
	}

	index.Metrics = filteredList
	return index, nil
}
