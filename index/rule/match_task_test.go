// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchTaskAddSentence(t *testing.T) {
	var matchTask = NewMatchTask()

	matchTask.AddSentence("someContextMarker1", "This is a test1.")
	matchTask.AddSentence("someContextMarker2", "This is a test2.")

	if !assert.NotEmpty(t, matchTask.sentenceByContextMarker, "Match task contains text sentences.") {
		t.Fatal("Sentences are not added to the match task.")
	}

	var sentence1 = matchTask.sentenceByContextMarker["someContextMarker1"]
	var sentence2 = matchTask.sentenceByContextMarker["someContextMarker2"]

	assert.NotEmpty(t, sentence1.words, "Sentence from match task doesn't contain words.")
	assert.NotEmpty(t, sentence2.words, "Sentence from match task doesn't contain words.")

	assert.Subset(
		t,
		sentence1.words,
		[...]string{"This", "is", "a", "test1."},
		"Sentence doesn't contain words from original string.",
	)

	assert.Subset(
		t,
		sentence2.words,
		[...]string{"This", "is", "a", "test2."},
		"Sentence doesn't contain words from original string.",
	)
}

func TestMatchTaskSize(t *testing.T) {
	var matchTask = NewMatchTask()

	matchTask.AddSentence("someContextMarker", "This is a test text 1.")
	matchTask.AddSentence("someContextMarker2", "This is a test text 2.")

	var matchTaskSize = matchTask.Size()

	assert.Equal(t, 12, matchTaskSize)
}

func BenchmarkMatchTaskSize(b *testing.B) {
	var matchTask = NewMatchTask()

	matchTask.AddSentence("title", "This is a test title.")
	matchTask.AddSentence("description", "This is a test description.")

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		matchTask.Size()
	}
}
