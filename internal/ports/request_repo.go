package ports

import (
	"time"
)

type WorkerRepo interface {
	AsyncHTTP()
	Run(url string, method string) (requestDuration time.Duration, responseSize int)
}
