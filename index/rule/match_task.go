// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"
	"strings"
)

// Represent a sentence from text entry, used to divide a whole text into
// small pieces (with presaved word order) suitable for concurrent processing.
type Sentence struct {
	// Offset in words from start of the text within a single context entry.
	// e.g. for entry contextMarker->text {"description": "Test, description"}
	// we can create a Sentence{"offset": 0, "words": ["Test,"]}
	// and another one Sentence{"offset": 1, "text": ["Description"]}
	offset int

	// Part of divided text entry as a set of ordered words.
	words []string
}

// Describes a task for a rule processor (rules matching against a text).
type MatchTask struct {
	// Context represents prepared text sentences (see "analysis" package),
	// aggregated by their semantical category or "marker" (e.g. job title; job description).
	// Some rules may check a specific context marker, to be applicable
	// to the whole text, e.g. a word is expected to be in the job title only,
	// then such rule becomes "matched".
	sentenceByContextMarker map[string]Sentence
}

// Adds a new text sentence under specific context with zero word offset.
func (t MatchTask) AddSentence(contextMarker string, words []string) {
	var sentence = Sentence{0, words}

	t.sentenceByContextMarker[contextMarker] = sentence
}

// Adds a new text sentence under specific context with specified word offset.
// Practically used by splitters to divide a task into small derived parts.
func (t MatchTask) addSentenceWithOffset(contextMarker string, text string, offset int) {
	var words = strings.Fields(text)
	var sentence = Sentence{offset, words}

	t.sentenceByContextMarker[contextMarker] = sentence
}

func NewMatchTask() MatchTask {
	return MatchTask{
		sentenceByContextMarker: make(map[string]Sentence),
	}
}

var matchTaskKey taskKey

func NewMatchTaskContext(task MatchTask) context.Context {
	return context.WithValue(context.Background(), matchTaskKey, task)
}

func MatchTaskFromContext(context context.Context) (MatchTask, bool) {
	task, isMatchTask := context.Value(matchTaskKey).(MatchTask)

	return task, isMatchTask
}
