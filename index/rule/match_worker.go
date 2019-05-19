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
		panic("context: match task context misuse.")
	}

	// TODO: do work.
	_ = matchTask

	for _, notifyChannel := range w.channelsToNotify {
		notifyChannel <- NewOccurrenceFoundEvent()
	}
}

func NewMatchWorker() *MatchWorker {
	return &MatchWorker{}
}
