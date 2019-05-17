// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

type EventType uint8

const (
	// Whenever a word that belongs to rule found in the text.
	OccurrenceFoundEvent EventType = iota
)

var ruleEventTypes = [...]string{
	"OccurrenceFound",
}

// fmt.Stringer implementation for event type.
func (et EventType) String() string {
	if et > OccurrenceFoundEvent {
		panic("rule: undefined event type.")
	}

	return ruleEventTypes[et]
}

type Events []Event

// Represents a rule event during the indexing process;
// for example, when a word that belongs to some rule was found by a worker
// and we need to modify index value according to the rule specification.
type Event struct {
	// The event type to decide receiver's business logic.
	t EventType

	// A set of rules that has been affected by the event.
	rules Rules
}

func NewOccurrenceFoundEvent() *Event {
	// TODO

	return &Event{OccurrenceFoundEvent, Rules{}}
}
