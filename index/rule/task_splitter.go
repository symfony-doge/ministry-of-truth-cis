// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
	"fmt"
	"sync"
)

var taskSplitterInstance *TaskSplitter

var taskSplitterOnce sync.Once

type UndefinedTaskSplitterError struct {
	task interface{}
}

// Implements error interface.
func (err UndefinedTaskSplitterError) Error() string {
	return fmt.Sprintf("Splitter for task is not defined (task=%T)", err.task)
}

// Represents a splitter for concurrent tasks;
// Defines algorithms how each task should be splitted to separate
// and independent parts for parallel execution.
type TaskSplitter struct {
	matchTaskSplitter *MatchTaskSplitter
}

// isSplittable returns positive if task can be divided into small subtasks.
func (s *TaskSplitter) isSplittable(task interface{}) (bool, error) {
	switch v := task.(type) {
	case MatchTask:
		return s.matchTaskSplitter.isSplittable(v)
	default:
		return false, UndefinedTaskSplitterError{task}
	}
}

// Returns a set of subtasks created from a given task.
func (s *TaskSplitter) Split(task interface{}, partsNum int) ([]context.Context, error) {
	switch v := task.(type) {
	case MatchTask:
		return s.matchTaskSplitter.Split(v, partsNum)
	default:
		return nil, UndefinedTaskSplitterError{task}
	}
}

func NewTaskSplitter() *TaskSplitter {
	return &TaskSplitter{
		matchTaskSplitter: NewMatchTaskSplitter(),
	}
}

func TaskSplitterInstance() *TaskSplitter {
	taskSplitterOnce.Do(func() {
		taskSplitterInstance = NewTaskSplitter()
	})

	return taskSplitterInstance
}
