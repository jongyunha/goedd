package application

import (
	"context"
	"goedd/internal/ddd"
	"goedd/stores/internal/application/commands"
	"goedd/stores/internal/application/queries"
	"goedd/stores/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
	}

	Queries interface {
		GetStore(ctx context.Context, query queries.GetStore) (*domain.MallStore, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CreateStoreHandler
	}

	appQueries struct {
		queries.GetStoreHandler
	}
)

var _ App = (*Application)(nil)

func New(
	stores domain.StoreRepository,
	catalog domain.CatalogRepository,
	mall domain.MallRepository,
	publisher ddd.EventPublisher[ddd.Event],
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler: commands.NewCreateStoreHandler(stores, publisher),
		},
		appQueries: appQueries{
			GetStoreHandler: queries.NewGetStoreHandler(mall),
		},
	}
}
