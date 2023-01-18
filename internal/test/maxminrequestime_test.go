package test

import (
	"load-test/internal/utils"
	"testing"
	"time"
)

func Test_FindMaxRequestTime(t *testing.T) {
	tm := func(d string) time.Duration {
		dr, _ := time.ParseDuration(d)
		return dr
	}
	type args struct {
		t1 time.Duration
		t2 time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "test1",
			args: args{
				t1: tm("10s"),
				t2: tm("5s"),
			},
			want: tm("10s"),
		},
		{
			name: "test2",
			args: args{
				t1: tm("5s"),
				t2: tm("10s"),
			},
			want: tm("10s"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.FindMaxRequestTime(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("findMaxRequestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinRequestTime(t *testing.T) {
	tm := func(d string) time.Duration {
		dr, _ := time.ParseDuration(d)
		return dr
	}
	type args struct {
		t1 time.Duration
		t2 time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "test1",
			args: args{
				t1: tm("10s"),
				t2: tm("5s"),
			},
			want: tm("5s"),
		},
		{
			name: "test2",
			args: args{
				t1: tm("5s"),
				t2: tm("10s"),
			},
			want: tm("5s"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.FindMinRequestTime(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("findMinRequestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
