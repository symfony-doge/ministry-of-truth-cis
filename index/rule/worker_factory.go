// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"fmt"
)

// Worker factory is responsible for worker instantiation.
// It receives a task instance and should return a valid worker.
type WorkerFactory interface {
	CreateFor(interface{}) (Worker, error)
}

type UndefinedWorkerError struct {
	task interface{}
}

// Implements error interface.
func (err UndefinedWorkerError) Error() string {
	return fmt.Sprintf("Worker for task is not defined (task=%T)", err.task)
}
