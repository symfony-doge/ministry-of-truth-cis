// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
)

// Represents splitter for match task.
type MatchTaskSplitter struct{}

// Splits task into partsCount separate tasks.
func (s *MatchTaskSplitter) Split(task MatchTask, partsCount int) ([]context.Context, error) {
	// A single execution flow case.
	if partsCount < 2 {
		return []context.Context{NewMatchTaskContext(task)}, nil
	}

	var contexts = make([]context.Context, partsCount)

	// TODO split algorithm.
	for contextNum := range contexts {
		contexts[contextNum] = NewMatchTaskContext(task)
	}

	return contexts, nil
}

func NewMatchTaskSplitter() *MatchTaskSplitter {
	return &MatchTaskSplitter{}
}
