// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
)

// Describes a task for a rule processor (rules matching against a text).
type MatchTask struct {
	// Context represents prepared text sentences (see "analysis" package),
	// aggregated by their semantical category or "marker" (e.g. job title; job description).
	// Some rules may check a specific context marker, to be applicable
	// to the whole text, e.g. a word is expected to be in the job title only,
	// then such rule becomes "matched".
	textByContextMarker map[string]string
}

// Adds a new text sentence under specific context.
func (t MatchTask) AddSentence(contextMarker string, text string) {
	t.textByContextMarker[contextMarker] = text
}

// Implements ConcurrentTask interface.
func (t MatchTask) Split(partsNum int) []context.Context {
	// TODO

	var parts = make([]context.Context, partsNum)

	for partNum := range parts {
		parts[partNum] = context.TODO()
	}

	return parts
}

func NewMatchTask() MatchTask {
	return MatchTask{
		textByContextMarker: make(map[string]string),
	}
}
