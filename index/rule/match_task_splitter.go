// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
)

// 1) При дизайне структуры заранее учитывать разделение на подзадачи
// (например, взять дерево, либо другую структуру, которую можно асинхронно
// строить и контекстно-независимо обходить за 0(1) или O(log N).
// 2) Метод разделения на подзадачи должен быть недорогим, помогает
// предварительно задизайненная под параллельные вычисления структура, см п.1
// 3) Определить порог (объем данных), до которого запускать задачу в одном
// потоке выполнения, иначе будет оверхед на коммуникацию.

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

// Represents splitter for match task.
type MatchTaskSplitter struct{}

// Splits task into partsCount separate tasks.
func (s *MatchTaskSplitter) Split(task MatchTask, partsCount int) ([]context.Context, error) {
	// A single execution flow case.
	if partsCount < 2 {
		return []context.Context{NewMatchTaskContext(task)}, nil
	}

	var splitContext = mtSplitContext{}
	splitContext.task = &task
	splitContext.partSize = s.calculatePartSize(task, partsCount)
	splitContext.upperBound = splitContext.partSize
	splitContext.contexts = make([]context.Context, partsCount)

	for contextMarker := range task.sentenceByContextMarker {
		splitContext.contextMarker = contextMarker

		for {
			if isEndOfContext := s.splitNext(&splitContext); isEndOfContext {
				break
			}
		}
	}

	return splitContext.contexts, nil
}

func (s *MatchTaskSplitter) calculatePartSize(task MatchTask, partsCount int) int {
	var lenSum = 0

	// TODO: cache len in add sentence call(?)
	for _, sentence := range task.sentenceByContextMarker {
		lenSum += len(sentence.words)
	}

	return lenSum / partsCount
}

// TODO refactoring is required.
func (s *MatchTaskSplitter) splitNext(splitContext *mtSplitContext) (isEndOfContext bool) {
	var sentence = splitContext.task.sentenceByContextMarker[splitContext.contextMarker]
	var wordsCount = len(sentence.words)
	var isLastPart = splitContext.currentContextIndex == len(splitContext.contexts)-1

	if splitContext.upperBound > wordsCount || isLastPart {
		splitContext.upperBound -= wordsCount
		isEndOfContext = true
	}

	{
		var workerContext = splitContext.contexts[splitContext.currentContextIndex]
		if nil == workerContext {
			var partialTask = NewMatchTask()
			partialTask.addSentenceWithOffset(splitContext.contextMarker, []string{}, splitContext.lowerBound)
			splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(partialTask)
		}
	}

	var workerContext = splitContext.contexts[splitContext.currentContextIndex]
	partialTask, _ := MatchTaskFromContext(workerContext)
	if isEndOfContext {
		partialTask.addWordsToSentence(splitContext.contextMarker, sentence.words[splitContext.lowerBound:])
		splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(partialTask)
		splitContext.lowerBound = 0
	} else {
		partialTask.addWordsToSentence(splitContext.contextMarker, sentence.words[splitContext.lowerBound:splitContext.upperBound])
		splitContext.contexts[splitContext.currentContextIndex] = NewMatchTaskContext(partialTask)
		splitContext.lowerBound = splitContext.upperBound
		splitContext.upperBound += splitContext.partSize
		splitContext.currentContextIndex += 1
	}

	return
}

func NewMatchTaskSplitter() *MatchTaskSplitter {
	return &MatchTaskSplitter{}
}
