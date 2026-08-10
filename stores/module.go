package stores

import (
	"context"
	"goedd/internal/di"
	"goedd/internal/system"
	"goedd/stores/internal/constants"
)

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
	})
}
