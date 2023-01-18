package commands

import (
	"load-test/internal/ports"
)

type (
	CreateRequest struct {
	}

	CreateRequestHandler struct {
		workerRepo ports.WorkerRepo
	}
)

// create Workerepo
func NewCreateStoreHandler(workeRepo ports.WorkerRepo) CreateRequestHandler {
	return CreateRequestHandler{
		workerRepo: workeRepo,
	}
}

// request to start load testing
func (a CreateRequestHandler) CreateRequest() {
	a.workerRepo.AsyncHTTP()

}
