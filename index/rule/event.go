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
	Type EventType

	// A set of rules that has been affected by the event.
	Rules

	// Additional data based on event type, e.g. OccurrenceFoundContext.
	Payload interface{}
}

// Payload for occurrence found event.
type OccurrenceFoundContext struct {
	word, contextMarker string
	offset              int
}

func NewOccurrenceFoundEvent(rules Rules, context OccurrenceFoundContext) Event {
	return NewEvent(OccurrenceFoundEvent, rules, context)
}

func NewEvent(t EventType, rules Rules, payload interface{}) Event {
	return Event{t, rules, payload}
}
