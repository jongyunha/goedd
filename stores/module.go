package stores

import (
	"context"
	"goedd/internal/ddd"
	"goedd/internal/di"
	"goedd/internal/jetstream"
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
	_ = jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	return
}

func registrations(reg registry.Registry) (err error) {
	return
}
