// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/config"
)

const (
	// Path to json file with rules.
	configPathRulesJson string = "data.rule.json"
)

var invertedIndexInstance *InvertedIndex

var invertedIndexOnce sync.Once

type rulesByWord map[string][]*Rule

type wordsByContext map[string]rulesByWord

// Represents a data structure in which words or separate letters are
// associated with rules. Provides O(log n) lookup method.
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
	wordsByContext
}

// TODO building & lookup api, initialization in main func at start.

func (i *InvertedIndex) Build() error {
	rules, loadErr := i.loadRules()
	if nil != loadErr {
		return loadErr
	}

	for ruleNum := range rules {
		i.addToIndex(rules[ruleNum])
	}

	return nil
}

func (i *InvertedIndex) loadRules() (Rules, error) {
	var c = config.Instance()
	var filename = c.GetString(configPathRulesJson)

	var buf, readErr = ioutil.ReadFile(filename)
	if nil != readErr {
		return nil, readErr
	}

	var rules Rules
	if unmarshalErr := json.Unmarshal(buf, &rules); nil != unmarshalErr {
		return nil, unmarshalErr
	}

	return rules, nil
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
		rbw[word] = []*Rule{rule}
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
	return &InvertedIndex{make(wordsByContext)}
}

func InvertedIndexInstance() *InvertedIndex {
	invertedIndexOnce.Do(func() {
		invertedIndexInstance = NewInvertedIndex()
	})

	return invertedIndexInstance
}
