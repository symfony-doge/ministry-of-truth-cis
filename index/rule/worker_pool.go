// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"fmt"
	"sync"
)

// Splits a task to separate parts and distributes their execution
// among all available workers; related task splitter and a factory method
// to construct appropriate worker should be defined.
type WorkerPool interface {
	// Receives a concurrent task and a channel for worker events.
	// Returns a wait group instance if workers are successfully started.
	Distribute(interface{}, chan<- Event) (*sync.WaitGroup, error)
}

type WorkerNotPreparedError struct {
	task interface{}
}

// Implements error interface.
func (err WorkerNotPreparedError) Error() string {
	return fmt.Sprintf("Unable to prepare workers for distribution (task=%T)", err.task)
}
