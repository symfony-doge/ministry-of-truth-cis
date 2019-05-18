// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"
)

var defaultWorkerFactoryInstance *DefaultWorkerFactory

var defaultWorkerFactoryOnce sync.Once

type DefaultWorkerFactory struct{}

func (wf *DefaultWorkerFactory) CreateFor(task ConcurrentTask) (Worker, error) {
	switch task.(type) {
	case MatchTask:
		return NewMatchWorker(), nil
	default:
		return nil, UndefinedTaskError{task}
	}
}

func NewDefaultWorkerFactory() *DefaultWorkerFactory {
	return &DefaultWorkerFactory{}
}

func DefaultWorkerFactoryInstance() *DefaultWorkerFactory {
	defaultWorkerFactoryOnce.Do(func() {
		defaultWorkerFactoryInstance = NewDefaultWorkerFactory()
	})

	return defaultWorkerFactoryInstance
}
