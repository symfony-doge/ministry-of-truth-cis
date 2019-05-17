// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
)

// Uses a worker pool and a listener to subscribe for rule events
// and collect results (rule occurrences) using all available CPU cores.
// Use constructor function to create a new instance.
type ConcurrentProcessor struct {
	logger *log.Logger

	// Acquires events from workers.
	listener EventListener
}

func (p *ConcurrentProcessor) FindMatch(text string) (Rules, error) {
	notifyChannel, err := p.listener.Listen(p.onRuleEvent)
	if nil != err {
		p.logger.Println(err)

		return nil, EventListenerNotStartedError{}
	}

	// TODO distribute tasks among workers.
	notifyChannel <- *NewOccurrenceFoundEvent()
	notifyChannel <- *NewOccurrenceFoundEvent()
	close(notifyChannel)

	return Rules{}, nil
}

// Fires each time when a new rule event is available for processing.
func (p *ConcurrentProcessor) onRuleEvent(event Event) {
	// TODO

	log.Printf("Event consumed by the processor: %v\n", event)
}

func NewConcurrentProcessor() *ConcurrentProcessor {
	return &ConcurrentProcessor{
		logger:   DefaultLogger,
		listener: DefaultEventListenerInstance(),
	}
}
