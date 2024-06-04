package test

import (
	"factorialfrontend/api"
	"factorialfrontend/domain"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewIndex(t *testing.T) {
	type args struct {
		api api.IApi[domain.Metric]
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Index
		wantErr bool
	}{
		{
			name: "TestNewIndex",
			args: args{
				api: MockApi[domain.Metric]{},
			},
			want:    domain.Index{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewIndex(tt.args.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndexFiltered(t *testing.T) {
	type args struct {
		api        api.IApi[domain.Metric]
		filterTime string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Index
		wantErr bool
	}{
		{
			name: "TestNewIndexFiltered",
			args: args{
				api:        MockApi[domain.Metric]{},
				filterTime: "hour",
			},
			want: domain.Index{
				Metrics:        []domain.Metric{},
				IsChartEnabled: true,
				ChartData: domain.Chart{
					Label:  "Last Day (Average 0.0 per hour)",
					Labels: []string{fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*24).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*23).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*22).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*21).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*20).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*19).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*18).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*17).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*16).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*15).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*14).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*13).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*12).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*11).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*10).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*9).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*8).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*7).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*6).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*5).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*4).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*3).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour*2).Hour()), fmt.Sprintf("%02d", time.Now().Add(-1*time.Hour).Hour()), fmt.Sprintf("%02d", time.Now().Hour())},
					Data:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			wantErr: false,
		},
		{
			name: "TestNewIndexFiltered_empty",
			args: args{
				api:        MockApi[domain.Metric]{},
				filterTime: "",
			},
			want: domain.Index{
				Metrics:        []domain.Metric{},
				IsChartEnabled: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewIndexFiltered(tt.args.api, tt.args.filterTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIndexFiltered() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(cmp.Equal(tt.want.ChartData, got.ChartData) && tt.want.IsChartEnabled == got.IsChartEnabled) {
				t.Errorf("NewIndexFiltered() = %v, want %v", got, tt.want)
			}
		})
	}
}
