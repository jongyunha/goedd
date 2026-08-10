package stores

import (
	"context"
	"goedd/internal/di"
	"goedd/internal/registry"
	"goedd/internal/system"
	"goedd/stores/internal/constants"
	"goedd/stores/storespb"
)

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})

	return
}

func registrations(reg registry.Registry) (err error) {
	return
}
