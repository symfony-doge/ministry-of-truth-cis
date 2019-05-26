// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

const (
	atTestLocale    request.Locale = "ru"
	atTestRulesJson string         = "../data/rules.json"
)

var (
	atTestTagNamesExpected = []string{"young_dynamic"}
)

func atMockedAggregator() *Aggregator {
	var b *Bag = btMockedBag() // bag_test.go
	b.Load(atTestLocale)

	var a *Aggregator = AggregatorInstance()
	a.tagBag = b

	return a
}

func TestAggregatorExtractTagNames(t *testing.T) {
	var a *Aggregator = atMockedAggregator()

	var ruleProvider = rule.JSONProviderInstance()
	var rules, loadErr = ruleProvider.GetRulesFrom(atTestRulesJson)
	if nil != loadErr {
		t.Fatal("Unable to load rules:", loadErr)
	}

	var tagNames = a.ExtractTagNames(rules)

	assert.Equal(
		t,
		atTestTagNamesExpected,
		tagNames,
		"Expecting tag names has extracted from rules.",
	)
}

func TestAggregatorAggregateByGroup(t *testing.T) {
	var a *Aggregator = atMockedAggregator()

	var tagsByGroups = a.AggregateByGroup(atTestTagNamesExpected, atTestLocale)

	assert.Contains(t, tagsByGroups, "soft", "The result map contains an expected group name.")

	var tags = tagsByGroups["soft"]

	if !assert.NotEmpty(t, tags, "Expecting tags slice is not empty.") {
		t.Fatal("Tags slice is empty (len >= 1 expected).")
	}

	assert.Equal(
		t,
		"Динамично развивающаяся компания",
		tags[0].Title,
		"Expecting a valid tag instance.",
	)
}
