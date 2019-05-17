// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// ConsumeFunc is a callback signature for managing received events.
// It receives a read-only copy of a rule event.
type ConsumeFunc func(Event)

// Listens rule events from workers
// and calls specified closure for processing.
type EventListener interface {
	Listen(ConsumeFunc) (chan<- Event, error)
}

// Listener session is not designed to be a thread-safe unit
// and should not be accessed from multiple execution flows, i.e.
// only listener can access and manage it, internally.
type EventListenerSession struct {
	// Holds all received events within a single listening session.
	received Events

	// Index of the next event to be "consumed",
	// i.e. processed with ConsumeFunc.
	consumed int

	// Receives a positive value when listening is done.
	done bool
}

func NewEventListenerSession() *EventListenerSession {
	return &EventListenerSession{}
}
