// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
)

// Represents a task for the concurrent rule processor; concurrent task
// should define the algorithm how it will be splitted to separate
// and independent parts for parallel execution.
type ConcurrentTask interface {
	Split(int) []context.Context
}
