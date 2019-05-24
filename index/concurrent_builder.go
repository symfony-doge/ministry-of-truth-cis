// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"log"

	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
)

// Uses parallel algorithms to build a sanity index.
type ConcurrentBuilder struct {
	logger *log.Logger

	// Returns a sanity index value by applicable rules.
	valueCalculator
}

func (b *ConcurrentBuilder) Build(context BuilderContext) *Index {
	var ruleMatchTask = rule.NewMatchTask()

	for contextMarker, text := range context {
		ruleMatchTask.AddSentence(contextMarker, text)
	}

	// Returns all sanity rules applicable to the specified context.
	var ruleProcessor rule.Processor = rule.NewConcurrentProcessor()

	applicableRules, matchingErr := ruleProcessor.FindMatch(ruleMatchTask)
	if nil != matchingErr {
		b.logger.Println(matchingErr)

		panic("index: unable to find applicable rules for the text.")
	}

	var sanityIndex *Index = newIndex()

	sanityIndex.Value = b.valueCalculator.Calculate(applicableRules)

	return sanityIndex
}

func NewConcurrentBuilder() *ConcurrentBuilder {
	return &ConcurrentBuilder{
		logger:          DefaultLogger,
		valueCalculator: weightedAverageCalculatorInstance(),
	}
}
