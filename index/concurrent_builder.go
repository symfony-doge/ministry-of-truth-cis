// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"log"

	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
	"github.com/symfony-doge/ministry-of-truth-cis/tag"
)

// Uses parallel algorithms to build a sanity index.
type ConcurrentBuilder struct {
	logger *log.Logger

	// Returns a sanity index value by applicable rules.
	valueCalculator

	// Performs tags aggregation by group names.
	tagAggregator *tag.Aggregator
}

func (b *ConcurrentBuilder) Build(context BuilderContext, locale request.Locale) *Index {
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

	var tagNames = b.tagAggregator.ExtractTagNames(applicableRules)
	sanityIndex.Tags = b.tagAggregator.AggregateByGroup(tagNames, locale)

	return sanityIndex
}

func NewConcurrentBuilder() *ConcurrentBuilder {
	return &ConcurrentBuilder{
		logger:          DefaultLogger,
		valueCalculator: weightedAverageCalculatorInstance(),
		tagAggregator:   tag.AggregatorInstance(),
	}
}
