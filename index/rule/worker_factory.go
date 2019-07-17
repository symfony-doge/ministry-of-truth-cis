// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"

	"github.com/symfony-doge/splitex"
)

var workerFactoryInstance *WorkerFactory

var workerFactoryOnce sync.Once

// Implements splitex.WorkerFactory interface.
type WorkerFactory struct{}

func (wf *WorkerFactory) CreateFor(task interface{}) (splitex.Worker, error) {
	switch task.(type) {
	case MatchTask:
		return NewMatchWorker(), nil
	default:
		return nil, splitex.UndefinedWorkerError{task}
	}
}

func NewWorkerFactory() *WorkerFactory {
	return &WorkerFactory{}
}

func WorkerFactoryInstance() *WorkerFactory {
	workerFactoryOnce.Do(func() {
		workerFactoryInstance = NewWorkerFactory()
	})

	return workerFactoryInstance
}
