// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"
)

// Splits a task to separate parts and distributes their execution
// among all available workers. Task must implement ConcurrentTask interface
// with Split method.
type WorkerPool interface {
	// Sets a worker factory that creates workers for specific tasks.
	SetWorkerFactory(WorkerFactory)

	// Receives an input text and a channel for worker events.
	// Returns a wait group instance if workers are successfully started.
	Distribute(ConcurrentTask, chan<- Event) (*sync.WaitGroup, error)
}
