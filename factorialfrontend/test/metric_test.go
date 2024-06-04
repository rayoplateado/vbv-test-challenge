package test

import (
	"factorialfrontend/domain"
	"reflect"
	"testing"
)

func TestMetric_List(t *testing.T) {
	tests := []struct {
		name    string
		metric  domain.Metric
		want    []domain.Metric
		wantErr bool
	}{
		{
			name:    "TestMetric_List",
			metric:  domain.Metric{API: MockApi[domain.Metric]{}},
			want:    []domain.Metric{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.metric.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && (len(got) != len(tt.want) && len(got) == 0) {
				t.Errorf("Metric.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetric_FindOne(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		metric  domain.Metric
		args    args
		want    domain.Metric
		wantErr bool
	}{
		{
			name:   "TestMetric_FindOne",
			metric: domain.Metric{API: MockApi[domain.Metric]{}},
			args: args{
				id: "1",
			},
			want:    domain.Metric{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.metric.FindOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metric.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetric_Add(t *testing.T) {
	tests := []struct {
		name    string
		metric  domain.Metric
		want    domain.Metric
		wantErr bool
	}{
		{
			name:    "TestMetric_Add",
			metric:  domain.Metric{API: MockApi[domain.Metric]{}},
			want:    domain.Metric{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.metric.Add()
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metric.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
