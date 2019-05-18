// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
	"runtime"
	"sync"
)

const (
	// Number of reserved execution flows (for environment with 2 and more CPU),
	// e.g. for listening and processing results from workers simultaneously.
	executionFlowsReserved = 1
)

// Creates a worker for each available CPU, that can be executing simultaneously,
// and makes them process a part of specified task.
type DefaultWorkerPool struct {
	workerFactory WorkerFactory

	workers []Worker
}

func (wp *DefaultWorkerPool) SetWorkerFactory(wf WorkerFactory) {
	wp.workerFactory = wf
}

func (wp *DefaultWorkerPool) Distribute(
	task ConcurrentTask,
	notifyChannel chan<- Event,
) (*sync.WaitGroup, error) {
	if err := wp.prepareWorkers(task, notifyChannel); nil != err {
		return nil, err
	}

	return wp.start()
}

// Creates workers and sets their execution contexts.
func (wp *DefaultWorkerPool) prepareWorkers(
	task ConcurrentTask,
	notifyChannel chan<- Event,
) error {
	var workerCount int = runtime.GOMAXPROCS(0) - executionFlowsReserved
	if workerCount < 1 {
		workerCount = 1
	}

	wp.workers = make([]Worker, workerCount)
	var contexts []context.Context = task.Split(workerCount)

	for workerNumber := range wp.workers {
		var worker, wfErr = wp.workerFactory.CreateFor(task)
		if nil != wfErr {
			return wfErr
		}

		worker.SetContext(contexts[workerNumber])
		worker.AddNotifyChannel(notifyChannel)

		wp.workers[workerNumber] = worker
	}

	return nil
}

// Runs all prepared workers and returns a wait group to directly
// track their activity.
func (wp *DefaultWorkerPool) start() (*sync.WaitGroup, error) {
	var workerCount = len(wp.workers)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workerCount)

	for workerNumber := range wp.workers {
		go func() {
			defer waitGroup.Done()

			wp.workers[workerNumber].Run()
		}()
	}

	return &waitGroup, nil
}

func NewDefaultWorkerPool() *DefaultWorkerPool {
	return &DefaultWorkerPool{
		workerFactory: DefaultWorkerFactoryInstance(),
	}
}
