// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/analysis"
)

var invertedOccurrenceFinderInstance *InvertedOccurrenceFinder

var invertedOccurrenceFinderOnce sync.Once

// A component that uses an inverted index data structure to search
// stop-words and check when a rule is applicable to the text or not.
type InvertedOccurrenceFinder struct {
	logger *log.Logger

	// Transforms a word from text to its stem.
	wordStemmer analysis.Stemmer

	// Performs rules matching by stemmed word.
	invertedIndex *InvertedIndex
}

func (of *InvertedOccurrenceFinder) FindApplicableRules(word, contextMarker string) (Rules, bool) {
	wordStemmed, stemmingErr := of.wordStemmer.Stem(word)
	if nil != stemmingErr {
		of.logger.Println(stemmingErr)

		panic("index: unable to stem a text.")
	}

	return of.invertedIndex.Lookup(wordStemmed, contextMarker)
}

func NewInvertedOccurrenceFinder() *InvertedOccurrenceFinder {
	return &InvertedOccurrenceFinder{
		logger:        DefaultLogger,
		wordStemmer:   analysis.SnowballStemmerInstance(),
		invertedIndex: InvertedIndexInstance(),
	}
}

func InvertedOccurrenceFinderInstance() *InvertedOccurrenceFinder {
	invertedOccurrenceFinderOnce.Do(func() {
		invertedOccurrenceFinderInstance = NewInvertedOccurrenceFinder()
	})

	return invertedOccurrenceFinderInstance
}
