// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// ConsumeFunc is a callback signature for managing received events.
// It receives a read-only copy of a rule event from listening session.
type ConsumeFunc func(Event)

// Listens rule events from workers and calls a specified closure for processing.
// Returns a read-only listening session instance that exposes a communication
// channel between workers and event listener.
//
// Usage example:
//
// listenerSession, listenErr := eventListener.Listen(func(e Event) {})
// if listenErr != nil {
//     // Handle error.
// }
// var notifyChannel chan<- Event = listenerSession.NotifyChannel()
// notifyChannel <- Event{}
// listenerSession.Close()    // Use only this method to safely close the channel.
//
type EventListener interface {
	Listen(ConsumeFunc) (ROEventListenerSession, error)
}

// The event listener starts one goroutine per listening session.
// Session is not designed to be a thread-safe unit
// and should not be accessed from multiple execution flows, i.e.
// only listener and internals can access and manage it.
type eventListenerSession struct {
	// Holds all received events within a single listening session.
	received Events

	// Index of the next event to be "consumed", i.e. processed with ConsumeFunc.
	consumed int

	// Receives a value when listening is done.
	done chan bool
}

// Read-only listener session is used as a bridge between event listener
// and workers; a caller should pass the notification channel associated
// with listening session to the senders to start communication via events.
type ROEventListenerSession struct {
	els *eventListenerSession

	// A channel that should be used by workers to push their events.
	// If a notification channel becomes closed, listening session ends.
	// Notification channel must be closed only by Close() method, it is not
	// safe to use a built-in close function directly due to non-blocking
	// event processing.
	notifyChannel chan<- Event
}

// Returns a notification channel for sending events.
func (ls ROEventListenerSession) NotifyChannel() chan<- Event {
	return ls.notifyChannel
}

// Use it to safely close the notification channel and stop event listening session.
func (ls ROEventListenerSession) Close() {
	// This will ensure that all remaining events are properly processed.
	defer ls.wait()

	close(ls.notifyChannel)
}

func (ls ROEventListenerSession) wait() {
	<-ls.els.done
}

func newEventListenerSession() *eventListenerSession {
	return &eventListenerSession{
		received: Events{},
		consumed: 0,
		done:     make(chan bool, 1),
	}
}

func newROEventListenerSession(els *eventListenerSession, nc chan<- Event) ROEventListenerSession {
	return ROEventListenerSession{els, nc}
}
