package domain

import (
	"time"
)

// Status shows the current api status.
type Status struct {
	TotalDuration     time.Duration // total duration is to calculate the time for all request
	MinTime           time.Duration // for finding the fastest request
	MaxTime           time.Duration //  for finding the slowest request
	RequestsCounter   int           // counting the request
	TotalResponseSize int64         // size of the response body + header
	AverageTime       time.Duration // calculate the average time for all request
	ErrorCount        int           // if we don't get the response  from the request,. we are counting error base on the error
}

// new status
func NewStatus() Status {
	return Status{
		TotalDuration: time.Millisecond,
		MinTime:       time.Millisecond,
		MaxTime:       time.Millisecond,
	}
}
