// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"log"

	"github.com/symfony-doge/ministry-of-truth-cis/analysis"
	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
)

// Uses parallel algorithms to build a sanity index.
type ConcurrentBuilder struct {
	logger *log.Logger

	// Transforms all words in text to their initial form.
	textLemmatizer analysis.Lemmatizer

	// Returns all sanity rules applicable to the specified context.
	ruleProcessor rule.Processor
}

func (b *ConcurrentBuilder) Build(context BuilderContext) *Index {
	var description string = context[BuilderContextDescription]

	descriptionLemmatized, lErr := b.textLemmatizer.Lemmatize(description)
	if nil != lErr {
		b.logger.Println(lErr)

		panic("index: unable to lemmatize text.")
	}

	applicableRules, bErr := b.ruleProcessor.FindMatch(descriptionLemmatized)
	if nil != bErr {
		b.logger.Println(bErr)

		panic("index: unable to find applicable rules for the text.")
	}

	// TODO: index calculation by rules specifications.
	_ = applicableRules

	return NewIndex()
}

func NewConcurrentBuilder() *ConcurrentBuilder {
	return &ConcurrentBuilder{
		logger:         DefaultLogger,
		textLemmatizer: analysis.MyStemLemmatizerInstance(),
		ruleProcessor:  rule.NewConcurrentProcessor(),
	}
}
