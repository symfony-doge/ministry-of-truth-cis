// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"
)

var invertedOccurrenceFinderInstance *InvertedOccurrenceFinder

var invertedOccurrenceFinderOnce sync.Once

// A component that uses an inverted index data structure to search
// stop-words and check when a rule is applicable to the text or not.
type InvertedOccurrenceFinder struct{}

func (of *InvertedOccurrenceFinder) FindApplicableRules(word, contextMarker string) Rules {
	// TODO: search algorithm.

	return Rules{}
}

func NewInvertedOccurrenceFinder() *InvertedOccurrenceFinder {
	return &InvertedOccurrenceFinder{}
}

func InvertedOccurrenceFinderInstance() *InvertedOccurrenceFinder {
	invertedOccurrenceFinderOnce.Do(func() {
		invertedOccurrenceFinderInstance = NewInvertedOccurrenceFinder()
	})

	return invertedOccurrenceFinderInstance
}
