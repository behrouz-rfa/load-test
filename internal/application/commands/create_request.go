package commands

import (
	"load-test/internal/ports"
)

type (
	CreateRequest struct {
	}

	CreateRequestHandler struct {
		stores ports.WorkerRepo
	}
)

func NewCreateStoreHandler(stores ports.WorkerRepo) CreateRequestHandler {
	return CreateRequestHandler{
		stores: stores,
	}
}

func (a CreateRequestHandler) CreateRequest() {
	a.stores.AsyncHTTP()

}
