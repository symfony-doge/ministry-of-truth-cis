// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
	"runtime"
	"sync"

	"github.com/symfony-doge/event"
)

const (
	// Number of reserved execution flows (for environment with 2 and more CPU),
	// e.g. for listening and processing results from workers simultaneously.
	executionFlowsReserved = 1
)

// Creates a worker for each available CPU, that can be executing simultaneously,
// and makes them process a part of specified task.
type DefaultWorkerPool struct {
	logger *log.Logger

	// Encapsulates split algorithms for concurrent tasks.
	taskSplitter *TaskSplitter

	// Instantiates workers.
	workerFactory WorkerFactory

	workers []Worker
}

func (wp *DefaultWorkerPool) Distribute(
	task interface{},
	notifyChannel chan<- event.Event,
) (*sync.WaitGroup, error) {
	if err := wp.prepareWorkers(task, notifyChannel); nil != err {
		wp.logger.Println(err)

		return nil, WorkerNotPreparedError{task}
	}

	return wp.runWorkers()
}

// Creates workers and sets their execution contexts.
func (wp *DefaultWorkerPool) prepareWorkers(
	task interface{},
	notifyChannel chan<- event.Event,
) error {
	var workerCount, resolvingErr = wp.resolveWorkerCount(task)
	if nil != resolvingErr {
		return resolvingErr
	}

	wp.workers = make([]Worker, workerCount)

	var contexts, splitErr = wp.taskSplitter.Split(task, workerCount)
	if nil != splitErr {
		return splitErr
	}

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

func (wp *DefaultWorkerPool) resolveWorkerCount(task interface{}) (int, error) {
	var workerCount int = runtime.GOMAXPROCS(0) - executionFlowsReserved

	// There is no reason to gain a splitting and communication overhead,
	// if only one execution flow is available.
	if workerCount < 1 {
		return 1, nil
	}

	// There is also can be a set of task-specific conditions,
	// when it should be splitted and when not (e.g. small data amount).
	var isTaskSplittable, checkErr = wp.taskSplitter.isSplittable(task)
	if nil != checkErr {
		return 0, checkErr
	}

	if !isTaskSplittable {
		return 1, nil
	}

	return workerCount, nil
}

// Runs all prepared workers and returns a wait group to directly
// track their activity.
func (wp *DefaultWorkerPool) runWorkers() (*sync.WaitGroup, error) {
	var workerCount = len(wp.workers)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workerCount)

	for workerNumber := range wp.workers {
		// We should not capture loop variables in closure, goroutine will
		// see only last assigned value; instead, we pass a copy as an argument.
		go func(wn int) {
			defer waitGroup.Done()
			wp.workers[wn].Run()
		}(workerNumber)
	}

	return &waitGroup, nil
}

func NewDefaultWorkerPool() *DefaultWorkerPool {
	return &DefaultWorkerPool{
		logger:        DefaultLogger,
		taskSplitter:  TaskSplitterInstance(),
		workerFactory: DefaultWorkerFactoryInstance(),
	}
}
