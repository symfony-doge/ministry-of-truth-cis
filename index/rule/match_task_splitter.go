// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// 1) При дизайне структуры заранее учитывать разделение на подзадачи
// (например, взять дерево, либо другую структуру, которую можно асинхронно
// строить и контекстно-независимо обходить за 0(1) или O(log N).
// 2) Метод разделения на подзадачи должен быть недорогим, помогает
// предварительно задизайненная под параллельные вычисления структура, см п.1
// 3) Определить порог (объем данных), до которого запускать задачу в одном
// потоке выполнения, иначе будет оверхед на коммуникацию.

package rule

import (
	"context"
)

const (
	// Minimum words required for task splitting.
	minWordsForSplitThreshold = 100
)

type mtSplitContext struct {
	// Task for splitting.
	task *MatchTask

	// Context marker.
	contextMarker string

	// Upper/Lower bounds for words slicing.
	lowerBound, upperBound int

	// Words count for each partial task.
	partSize int

	// Index for current context in building.
	currentContextIndex int

	// Result contexts.
	contexts []context.Context
}

// Represents a splitter for match tasks; divides a task to a set
// of separate and independent tasks for parallel processing.
type MatchTaskSplitter struct{}

func (s *MatchTaskSplitter) isSplittable(task MatchTask) (bool, error) {
	return task.Size() >= minWordsForSplitThreshold, nil
}

// Splits task into partsCount separate tasks.
func (s *MatchTaskSplitter) Split(task MatchTask, partsCount int) ([]context.Context, error) {
	// A single execution flow case, no splitting actually required.
	if partsCount < 2 {
		return []context.Context{NewMatchTaskContext(task)}, nil
	}

	var splitContext *mtSplitContext = s.newSplitContext(task, partsCount)

	for contextMarker := range task.sentenceByContextMarker {
		splitContext.contextMarker = contextMarker

		for {
			if isEndOfContext := s.splitNext(splitContext); isEndOfContext {
				break
			}
		}
	}

	return splitContext.contexts, nil
}

// Returns new context for task splitting operation.
func (s *MatchTaskSplitter) newSplitContext(task MatchTask, partsCount int) *mtSplitContext {
	var splitContext = &mtSplitContext{}

	splitContext.task = &task
	splitContext.partSize = s.calculatePartSize(task, partsCount)
	splitContext.upperBound = splitContext.partSize
	splitContext.contexts = make([]context.Context, partsCount)

	return splitContext
}

// Returns words count for processing by a single worker.
func (s *MatchTaskSplitter) calculatePartSize(task MatchTask, partsCount int) int {
	return task.Size() / partsCount
}

func (s *MatchTaskSplitter) splitNext(splitContext *mtSplitContext) (isEndOfContext bool) {
	var sentence = splitContext.task.sentenceByContextMarker[splitContext.contextMarker]
	var wordsCount = len(sentence.words)
	var isLastPart = splitContext.currentContextIndex == len(splitContext.contexts)-1

	if splitContext.upperBound > wordsCount || isLastPart {
		splitContext.upperBound -= wordsCount
		isEndOfContext = true
	}

	var workerContext context.Context = s.extractWorkerContext(splitContext)
	partialTask, _ := MatchTaskFromContext(workerContext)

	if isEndOfContext {
		s.fill(splitContext, &partialTask, sentence.words)
	} else {
		s.fillAndShift(splitContext, &partialTask, sentence.words)
	}

	return
}

func (s *MatchTaskSplitter) extractWorkerContext(splitContext *mtSplitContext) context.Context {
	s.ensureWorkerContext(splitContext)

	return splitContext.contexts[splitContext.currentContextIndex]
}

// Ensures worker context exists and is a valid instance within task splitting context.
func (s *MatchTaskSplitter) ensureWorkerContext(splitContext *mtSplitContext) {
	var workerContext = splitContext.contexts[splitContext.currentContextIndex]

	if nil == workerContext {
		// Empty partial task for filling.
		var partialTask = NewMatchTask()
		partialTask.addSentenceWithOffset(splitContext.contextMarker, []string{}, splitContext.lowerBound)

		splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(partialTask)
	}
}

func (s *MatchTaskSplitter) fill(
	splitContext *mtSplitContext,
	partialTask *MatchTask,
	words []string,
) {
	var partialWords = words[splitContext.lowerBound:]
	partialTask.addWordsToSentence(splitContext.contextMarker, partialWords)

	splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(*partialTask)

	splitContext.lowerBound = 0
}

func (s *MatchTaskSplitter) fillAndShift(
	splitContext *mtSplitContext,
	partialTask *MatchTask,
	words []string,
) {
	var partialWords = words[splitContext.lowerBound:splitContext.upperBound]
	partialTask.addWordsToSentence(splitContext.contextMarker, partialWords)

	splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(*partialTask)

	splitContext.lowerBound = splitContext.upperBound
	splitContext.upperBound += splitContext.partSize
	splitContext.currentContextIndex += 1
}

func NewMatchTaskSplitter() *MatchTaskSplitter {
	return &MatchTaskSplitter{}
}
