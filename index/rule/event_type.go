// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"github.com/symfony-doge/event"
)

const (
	// Whenever a word that belongs to rule found in the text.
	OccurrenceFoundEvent event.EventType = iota
)

// Payload for occurrence found event.
type OccurrenceFoundContext struct {
	// A set of rules that has been affected by the event.
	rules               Rules
	word, wordProcessed string
	offset              int
}

func NewOccurrenceFoundEvent(context OccurrenceFoundContext) event.Event {
	return event.WithTypeAndPayload(OccurrenceFoundEvent, context)
}
