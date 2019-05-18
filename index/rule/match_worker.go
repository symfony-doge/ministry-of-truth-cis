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
	partialMatchTask context.Context

	channelsToNotify []chan<- Event
}

func (w *MatchWorker) SetContext(context context.Context) {
	w.partialMatchTask = context
}

func (w *MatchWorker) AddNotifyChannel(notifyChannels ...chan<- Event) {
	w.channelsToNotify = append(w.channelsToNotify, notifyChannels...)
}

func (w *MatchWorker) Run() {
	// TODO

	for _, notifyChannel := range w.channelsToNotify {
		notifyChannel <- NewOccurrenceFoundEvent()
	}
}

func NewMatchWorker() *MatchWorker {
	return &MatchWorker{}
}
