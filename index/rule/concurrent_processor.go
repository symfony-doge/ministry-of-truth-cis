// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
)

// Can be thrown by a processor that uses a rule events system;
// indicates that an event listener has not been initialized.
type EventListenerNotStartedError struct{}

func (err EventListenerNotStartedError) Error() string {
	return "Unable to FindMatch. Event listener is not started."
}

// Can be thrown if workers for parallel execution is not started.
type WorkersNotStartedError struct{}

func (err WorkersNotStartedError) Error() string {
	return "Unable to FindMatch. Workers are not started."
}

// Uses a worker pool and a listener to subscribe for rule events
// and collect results (rule occurrences) using all available CPU cores.
// Use the construct function to create a new instance for each request.
type ConcurrentProcessor struct {
	logger *log.Logger

	// Splits a task to separate parts and distributes their execution
	// among all available workers.
	workerPool WorkerPool

	// Acquires events from workers.
	listener EventListener
}

func (p *ConcurrentProcessor) FindMatch(task MatchTask) (Rules, error) {
	notifyChannel, lErr := p.listener.Listen(p.onRuleEvent)
	if nil != lErr {
		p.logger.Println(lErr)

		return nil, EventListenerNotStartedError{}
	}

	var workersWaitGroup, wpErr = p.workerPool.Distribute(task, notifyChannel)
	if nil != wpErr {
		p.logger.Println(wpErr)

		return nil, WorkersNotStartedError{}
	}

	// Waiting while workers do their parts of task.
	workersWaitGroup.Wait()

	// Stops listening for new events after all workers is complete.
	close(notifyChannel)

	// TODO: return merged results, ensure all merged
	// (may be check for listener session is required).

	return Rules{}, nil
}

// Fires each time when a new rule event is available for processing.
func (p *ConcurrentProcessor) onRuleEvent(event Event) {
	// TODO results merging

	log.Printf("Event consumed by the processor: %v\n", event)
}

func NewConcurrentProcessor() *ConcurrentProcessor {
	return &ConcurrentProcessor{
		logger:     DefaultLogger,
		workerPool: NewDefaultWorkerPool(),
		listener:   DefaultEventListenerInstance(),
	}
}
