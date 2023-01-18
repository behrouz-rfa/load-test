package application

import (
	"load-test/internal/application/commands"
	"load-test/internal/ports"
)

type (
	App interface {
		Commands
	}

	Commands interface {
		CreateRequest()
	}

	Application struct {
		appCommands
	}
	appCommands struct {
		commands.CreateRequestHandler
	}
)

var _ App = (*Application)(nil)

func New(stores ports.WorkerRepo) *Application {
	return &Application{
		appCommands: appCommands{
			CreateRequestHandler: commands.NewCreateStoreHandler(stores),
		},
	}
}
