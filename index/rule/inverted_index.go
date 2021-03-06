// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"
)

var invertedIndexInstance *InvertedIndex

var invertedIndexOnce sync.Once

// Maps word to related rules.
type rulesByWord map[string]Rules

// Maps contextMarker to map[word]rules
type wordsByContext map[string]rulesByWord

// Represents a data structure in which words or separate letters are
// associated with rules. Provides O(1) for lookup method.
// Example for context "description" and stop-sentence "this is a test".
// {
//     "description": {
//         "this": [rule1, rule2],
//         "is": [rule1, rule2],
//         "a": [rule1, rule2],
//         "test": [rule1, rule2]
//     }
// }
type InvertedIndex struct {
	ruleProvider *JSONProvider

	wordsByContext
}

// Checks if word is binded to rules or not within specified context.
func (i *InvertedIndex) Lookup(word, contextMarker string) (Rules, bool) {
	if _, isContextMarkerFound := i.wordsByContext[contextMarker]; !isContextMarkerFound {
		return nil, false
	}

	if _, isWordFound := i.wordsByContext[contextMarker][word]; !isWordFound {
		return nil, false
	}

	return i.wordsByContext[contextMarker][word], true
}

func (i *InvertedIndex) Build() error {
	rules, loadErr := i.ruleProvider.GetRules()
	if nil != loadErr {
		return loadErr
	}

	for ruleNum := range rules {
		i.addToIndex(rules[ruleNum])
	}

	return nil
}

func (i *InvertedIndex) addToIndex(rule *Rule) {
	for _, rse := range rule.Specification {
		var rulesByWords = i.buildRulesByWords(rse.Words, rule)

		for cmIdx := range rse.Contexts {
			i.appendWordsToContext(rse.Contexts[cmIdx], rulesByWords)
		}
	}
}

func (i *InvertedIndex) buildRulesByWords(words []string, rule *Rule) rulesByWord {
	var rbw = make(rulesByWord)

	for _, word := range words {
		rbw[word] = Rules{rule}
	}

	return rbw
}

func (i *InvertedIndex) appendWordsToContext(contextMarker string, rbw rulesByWord) {
	if _, isContextMarkerExists := i.wordsByContext[contextMarker]; isContextMarkerExists {
		i.mergeRulesByWord(i.wordsByContext[contextMarker], rbw)

		return
	}

	i.wordsByContext[contextMarker] = rbw
}

func (i *InvertedIndex) mergeRulesByWord(to rulesByWord, from rulesByWord) {
	for word := range from {
		if _, isWordExists := to[word]; isWordExists {
			for ruleIdx := range from[word] {
				to[word] = append(to[word], from[word][ruleIdx])
			}

			continue
		}

		to[word] = from[word]
	}
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		ruleProvider:   JSONProviderInstance(),
		wordsByContext: make(wordsByContext),
	}
}

func InvertedIndexInstance() *InvertedIndex {
	invertedIndexOnce.Do(func() {
		invertedIndexInstance = NewInvertedIndex()
	})

	return invertedIndexInstance
}
