package domain

import (
	"fmt"
	"time"
)

type Chart struct {
	Labels []string
	Label  string
	Data   []int
}

func NewChart(filterTime string, metrics []Metric) Chart {
	var chartData = Chart{}
	var labels []string
	var data []int
	start := time.Now()
	count := 0

	switch filterTime {
	case "minute":
		start = start.Add(-1 * time.Hour)
		for {
			labels = append(labels, fmt.Sprintf("%02d", start.Minute()))
			filteredData := filter(metrics, func(metric Metric) bool {
				return metric.Timestamp.After(start) && metric.Timestamp.Before(start.Add(1*time.Minute))
			})
			start = start.Add(1 * time.Minute)
			data = append(data, len(filteredData))
			count = count + len(filteredData)
			if start.After(time.Now()) {
				break
			}
		}
		chartData.Label = "Last Hour (Average " + fmt.Sprintf("%.1f", float64(count)/float64(len(labels))) + " per minute)"
	case "hour":
		start = start.Add(-1 * time.Hour * 24)
		for {
			labels = append(labels, fmt.Sprintf("%02d", start.Hour()))
			filteredData := filter(metrics, func(metric Metric) bool {
				return metric.Timestamp.After(start) && metric.Timestamp.Before(start.Add(1*time.Hour))
			})
			start = start.Add(1 * time.Hour)
			data = append(data, len(filteredData))
			count = count + len(filteredData)
			if start.After(time.Now()) {
				break
			}
		}
		chartData.Label = "Last Day (Average " + fmt.Sprintf("%.1f", float64(count)/float64(len(labels))) + " per hour)"
	case "day":
		start = start.Add(-1 * time.Hour * 24 * 30)
		for {
			if start.After(time.Now()) {
				break
			}
			labels = append(labels, fmt.Sprintf("%02d", start.Day()))
			filteredData := filter(metrics, func(metric Metric) bool {
				return metric.Timestamp.After(start) && metric.Timestamp.Before(start.Add(1*time.Hour*24))
			})
			start = start.Add(1 * time.Hour * 24)
			data = append(data, len(filteredData))
			count = count + len(filteredData)
		}
		chartData.Label = "Last Month (Average " + fmt.Sprintf("%.1f", float64(count)/float64(len(labels))) + " per day)"
	}

	chartData.Data = data
	chartData.Labels = labels

	return chartData
}

func filter[T any](list []T, f func(param T) bool) []T {
	var returnList []T
	for _, t := range list {
		if f(t) {
			returnList = append(returnList, t)
		}
	}

	return returnList
}
