// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"fmt"
)

// Worker factory is responsible for worker instantiation.
type WorkerFactory interface {
	CreateFor(ConcurrentTask) (Worker, error)
}

type UndefinedTaskError struct {
	task ConcurrentTask
}

// Implements error interface.
func (err UndefinedTaskError) Error() string {
	return fmt.Sprintf("Worker for task is not defined (task=%T)", err.task)
}
