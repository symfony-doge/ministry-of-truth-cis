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
type InvertedOccurrenceFinder struct {
	invertedIndex *InvertedIndex
}

func (of *InvertedOccurrenceFinder) FindApplicableRules(word, contextMarker string) (Rules, bool) {
	return of.invertedIndex.Lookup(word, contextMarker)
}

func NewInvertedOccurrenceFinder() *InvertedOccurrenceFinder {
	return &InvertedOccurrenceFinder{
		invertedIndex: InvertedIndexInstance(),
	}
}

func InvertedOccurrenceFinderInstance() *InvertedOccurrenceFinder {
	invertedOccurrenceFinderOnce.Do(func() {
		invertedOccurrenceFinderInstance = NewInvertedOccurrenceFinder()
	})

	return invertedOccurrenceFinderInstance
}
