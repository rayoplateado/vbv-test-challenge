package test

import (
	"factorialapi/domain"
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
			metric:  domain.Metric{Collection: MockCollection[domain.Metric]{}},
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

func TestMetric_Find(t *testing.T) {
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
			name:   "TestMetric_Find",
			metric: domain.Metric{Collection: MockCollection[domain.Metric]{}},
			args: args{
				id: "1",
			},
			want:    domain.Metric{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.metric.Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metric.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetric_FindList(t *testing.T) {
	tests := []struct {
		name    string
		metric  domain.Metric
		want    []domain.Metric
		wantErr bool
	}{
		{
			name:    "TestMetric_FindList",
			metric:  domain.Metric{Collection: MockCollection[domain.Metric]{}},
			want:    []domain.Metric{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.metric.FindList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.FindList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && (len(got) != len(tt.want) && len(got) == 0) {
				t.Errorf("Metric.FindList() = %v, want %v", got, tt.want)
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
			metric:  domain.Metric{Collection: MockCollection[domain.Metric]{}},
			want:    domain.Metric{Collection: MockCollection[domain.Metric]{}},
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
