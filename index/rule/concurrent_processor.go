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
type WorkerNotStartedError struct{}

func (err WorkerNotStartedError) Error() string {
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
	eventListener EventListener

	// Performs merging of partial results from workers.
	matchTaskResultMerger *MatchTaskResultMerger
}

func (p *ConcurrentProcessor) FindMatch(task MatchTask) (Rules, error) {
	listenerSession, listenErr := p.eventListener.Listen(p.onRuleEvent)
	if nil != listenErr {
		p.logger.Println(listenErr)

		return nil, EventListenerNotStartedError{}
	}

	var notifyChannel = listenerSession.NotifyChannel()
	var workersWaitGroup, wpErr = p.workerPool.Distribute(task, notifyChannel)
	if nil != wpErr {
		p.logger.Println(wpErr)

		return nil, WorkerNotStartedError{}
	}

	// Waiting while workers do their parts of task.
	workersWaitGroup.Wait()

	// Stops listening for new events after all workers is complete; waits
	// for remain events to be properly processed.
	listenerSession.Close()

	var rules = p.matchTaskResultMerger.GetResult()

	return rules, nil
}

// Fires each time when a new rule event is available for processing.
// It is a result collecting/merging function for separate task parts.
func (p *ConcurrentProcessor) onRuleEvent(event Event) {
	switch event.Type {
	// Merging match task partial results.
	case OccurrenceFoundEvent:
		for ruleIdx := range event.Rules {
			var context, isOccurrenceFoundContext = event.Payload.(OccurrenceFoundContext)
			if !isOccurrenceFoundContext {
				panic("rule: occurrence found event misuse.")
			}

			p.matchTaskResultMerger.Merge(event.Rules[ruleIdx], context)
		}
	default:
		panic("rule: undefined event.")
	}
}

func NewConcurrentProcessor() *ConcurrentProcessor {
	return &ConcurrentProcessor{
		logger:                DefaultLogger,
		workerPool:            NewDefaultWorkerPool(),
		eventListener:         DefaultEventListenerInstance(),
		matchTaskResultMerger: NewMatchTaskResultMerger(),
	}
}
