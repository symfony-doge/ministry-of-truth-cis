// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
)

// Performs rule matching routine against a part of text sentence
// using settings from the specified context.
type MatchWorker struct {
	// Context with match task and any additional settings for worker.
	context context.Context

	channelsToNotify []chan<- Event

	// Returns rules which are applicable for the given word.
	occurrenceFinder OccurrenceFinder
}

func (w *MatchWorker) SetContext(context context.Context) {
	w.context = context
}

func (w *MatchWorker) AddNotifyChannel(notifyChannels ...chan<- Event) {
	w.channelsToNotify = append(w.channelsToNotify, notifyChannels...)
}

func (w *MatchWorker) Run() {
	var matchTask, isMatchTask = MatchTaskFromContext(w.context)
	if !isMatchTask {
		panic("rule: match task context misuse.")
	}

	for contextMarker, sentence := range matchTask.sentenceByContextMarker {
		w.checkWordOccurrence(contextMarker, sentence)
	}
}

func (w *MatchWorker) checkWordOccurrence(contextMarker string, sentence Sentence) {
	for wordOffset, word := range sentence.words {
		rules, isOccurrenceFound, wp := w.occurrenceFinder.FindApplicableRules(word, contextMarker)

		if !isOccurrenceFound {
			continue
		}

		var summaryOffset = sentence.offset + wordOffset
		var context = OccurrenceFoundContext{word, wp, summaryOffset}
		var occurrenceFoundEvent = NewOccurrenceFoundEvent(rules, context)

		for ncIdx := range w.channelsToNotify {
			w.channelsToNotify[ncIdx] <- occurrenceFoundEvent
		}
	}
}

func NewMatchWorker() *MatchWorker {
	return &MatchWorker{
		occurrenceFinder: InvertedOccurrenceFinderInstance(),
	}
}
