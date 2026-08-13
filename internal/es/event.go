package es

import "goedd/internal/ddd"

type EventApplier interface {
	ApplyEvent(event ddd.Event) error
}

type EventCommitter interface {
	CommitEvents()
}

func LoadEvent(v any, event ddd.AggregateEvent) error {
	type loader interface {
		EventApplier
		VersionSetter
	}

	return nil
}
