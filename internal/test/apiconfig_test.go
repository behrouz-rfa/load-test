package test

import (
	"load-test/internal/domain"
	"reflect"
	"testing"
)

func TestAPIConfig_stop(t *testing.T) {
	type fields struct {
		interrupt int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test1",
			fields: fields{
				interrupt: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &domain.APIConfig{
				Interrupt: tt.fields.interrupt,
			}
			conf.Stop()
		})
	}
}

func Test_newAPIConfig(t *testing.T) {
	type args struct {
		goroutines      int
		duration        int
		timeOut         int
		finalStatusChan chan *domain.Status
		params          *domain.RequestParams
	}
	tests := []struct {
		name string
		args args
		want *domain.APIConfig
	}{
		{
			name: "test1",
			args: args{
				goroutines:      10,
				duration:        10,
				timeOut:         1000,
				finalStatusChan: nil,
				params: &domain.RequestParams{
					URL:    "www.sample.com",
					Method: "POST",
					Header: nil,
				},
			},
			want: &domain.APIConfig{
				ConcurrentConnections: 10,
				Duration:              10,
				TimeOut:               1000,
				FinalStatus:           nil,
				Interrupt:             0,
				Params: &domain.RequestParams{

					URL:    "www.sample.com",
					Method: "POST",
					Header: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.NewAPIConfig(tt.args.goroutines, tt.args.duration, tt.args.timeOut, tt.args.finalStatusChan, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAPIConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
