// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"
)

const (
	// Max session contexts for default listener's history field.
	maxHistoryEntries = 10

	// Buffer size for events channel.
	// In ideal case we should not block workers during their communication
	// with listener, but it depends on concrete task and how often they
	// will communicate. This value is the medium.
	notifyChannelBufferSize = 1 << 3
)

var defaultEventListenerInstance *DefaultEventListener

var defaultEventListenerOnce sync.Once

// Implements a common event listening logic and supports
// multiple listening sessions.
type DefaultEventListener struct {
	// Listening history.
	history []*EventListenerSession
}

// Starts new listening session and returns a channel to which
// senders should push their events; stops listening when the notify channel
// becomes closed.
func (l *DefaultEventListener) Listen(fn ConsumeFunc) (chan<- Event, error) {
	var notifyChannel = make(chan Event, notifyChannelBufferSize)
	var listenerSession = NewEventListenerSession()

	go func() {
		for {
			select {
			// Blocks default flow whenever a new event is available to consume
			// or channel becomes closed.
			case event, isChannelOpen := <-notifyChannel:
				if !isChannelOpen {
					// We should set channel to nil, to ensure it will not block
					// a goroutine with infinity communication loop
					// (closed channels blocks immediately).
					listenerSession.done, notifyChannel = true, nil

					break
				}

				listenerSession.received = append(listenerSession.received, event)
			// While waiting for new events, we use goroutine time
			// to process ones which already received (non-blocking approach).
			default:
				// We still have to process events until each will be
				// "consumed", then we can check session state, but notify
				// channel is safe to be closed earlier.
				if listenerSession.consumed < len(listenerSession.received) {
					var next Event = listenerSession.received[listenerSession.consumed]

					fn(next)
					listenerSession.consumed++

					break
				}

				if listenerSession.done {
					return
				}
			}
		}
	}()

	l.rotateHistory(listenerSession)

	return notifyChannel, nil
}

func (l *DefaultEventListener) rotateHistory(els *EventListenerSession) {
	if len(l.history) >= maxHistoryEntries {
		l.history = make([]*EventListenerSession, maxHistoryEntries)
	}

	l.history = append(l.history, els)
}

func NewDefaultEventListener() *DefaultEventListener {
	return &DefaultEventListener{}
}

func DefaultEventListenerInstance() *DefaultEventListener {
	defaultEventListenerOnce.Do(func() {
		defaultEventListenerInstance = NewDefaultEventListener()
	})

	return defaultEventListenerInstance
}
