package test

import (
	"factorialfrontend/domain"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewChart(t *testing.T) {
	type args struct {
		filterTime string
		metrics    []domain.Metric
	}
	tests := []struct {
		name string
		args args
		want domain.Chart
	}{
		{
			name: "TestNewChart: Day, no results",
			args: args{
				filterTime: "day",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*30).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*29).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*28).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*27).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*26).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*25).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*24).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*23).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*22).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*21).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*20).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*19).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*18).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*17).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*16).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*15).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*14).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*13).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*12).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*11).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*10).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*9).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*8).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*7).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*6).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*5).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*4).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*3).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*2).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Day()), fmt.Sprintf("%02d", time.Now().Day())}, Label: "Last Month (Average 0.0 per day)", Data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Day, 1 result",
			args: args{
				filterTime: "day",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*30).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*29).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*28).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*27).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*26).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*25).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*24).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*23).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*22).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*21).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*20).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*19).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*18).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*17).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*16).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*15).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*14).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*13).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*12).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*11).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*10).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*9).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*8).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*7).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*6).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*5).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*4).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*3).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*2).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Day()), fmt.Sprintf("%02d", time.Now().Day())}, Label: "Last Month (Average 0.0 per day)", Data: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Day, 30 results",
			args: args{
				filterTime: "day",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 24 * 29)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*30).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*29).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*28).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*27).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*26).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*25).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*24).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*23).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*22).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*21).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*20).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*19).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*18).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*17).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*16).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*15).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*14).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*13).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*12).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*11).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*10).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*9).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*8).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*7).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*6).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*5).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*4).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*3).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24*2).Day()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Day()), fmt.Sprintf("%02d", time.Now().Day())}, Label: "Last Month (Average 1.0 per day)", Data: []int{30, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Hour, 0 result",
			args: args{
				filterTime: "hour",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*23).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*22).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*21).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*20).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*19).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*18).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*17).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*16).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*15).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*14).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*13).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*12).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*11).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*10).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*9).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*8).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*7).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*6).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*5).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*4).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*3).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*2).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour).Hour()), fmt.Sprintf("%02d", time.Now().Hour())}, Label: "Last Day (Average 0.0 per hour)", Data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Hour, 1 result",
			args: args{
				filterTime: "hour",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Hour * 23)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*23).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*22).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*21).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*20).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*19).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*18).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*17).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*16).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*15).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*14).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*13).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*12).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*11).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*10).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*9).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*8).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*7).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*6).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*5).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*4).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*3).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*2).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour).Hour()), fmt.Sprintf("%02d", time.Now().Hour())}, Label: "Last Day (Average 0.0 per hour)", Data: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Minute, 0 result",
			args: args{
				filterTime: "minute",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*60).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*59).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*58).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*57).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*56).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*55).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*54).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*53).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*52).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*51).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*50).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*49).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*48).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*47).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*46).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*45).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*44).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*43).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*42).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*41).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*40).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*39).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*38).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*37).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*36).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*35).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*34).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*33).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*32).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*31).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*30).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*29).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*28).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*27).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*26).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*25).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*24).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*23).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*22).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*21).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*20).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*19).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*18).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*17).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*16).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*15).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*14).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*13).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*12).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*11).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*10).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*9).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*8).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*7).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*6).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*5).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*4).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*3).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*2).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute).Minute()), fmt.Sprintf("%02d", time.Now().Minute())}, Label: "Last Hour (Average 0.0 per minute)", Data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "TestNewChart: Minute, 1 result",
			args: args{
				filterTime: "minute",
				metrics: []domain.Metric{
					{Name: "aaa", Value: "123", Timestamp: time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Date(2121, 1, 1, 1, 1, 1, 1, time.Local)},
					{Name: "aaa", Value: "123", Timestamp: time.Now().Add(-1 * time.Minute * 59)},
				},
			},
			want: domain.Chart{Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*60).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*59).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*58).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*57).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*56).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*55).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*54).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*53).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*52).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*51).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*50).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*49).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*48).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*47).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*46).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*45).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*44).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*43).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*42).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*41).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*40).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*39).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*38).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*37).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*36).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*35).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*34).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*33).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*32).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*31).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*30).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*29).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*28).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*27).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*26).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*25).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*24).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*23).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*22).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*21).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*20).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*19).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*18).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*17).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*16).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*15).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*14).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*13).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*12).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*11).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*10).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*9).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*8).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*7).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*6).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*5).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*4).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*3).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute*2).Minute()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Minute).Minute()), fmt.Sprintf("%02d", time.Now().Minute())}, Label: "Last Hour (Average 0.0 per minute)", Data: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.NewChart(tt.args.filterTime, tt.args.metrics); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChart() = %v, want %v", got, tt.want)
			}
		})
	}
}
