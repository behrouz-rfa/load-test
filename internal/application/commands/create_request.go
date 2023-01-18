package commands

import (
	"load-test/internal/ports"
)

type (
	CreateRequest struct {
	}

	CreateRequestHandler struct {
		woereRepo ports.WorkerRepo
	}
)

// create Workerepo
func NewCreateStoreHandler(workeRepo ports.WorkerRepo) CreateRequestHandler {
	return CreateRequestHandler{
		woereRepo: workeRepo,
	}
}

// request to start load testing
func (a CreateRequestHandler) CreateRequest() {
	a.woereRepo.AsyncHTTP()

}
