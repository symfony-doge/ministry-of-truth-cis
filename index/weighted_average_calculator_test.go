// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
)

const (
	wactIndexValueExpected = 95.5
)

func TestValueCalculatorWeightedAverageCalculate(t *testing.T) {
	var calculator = weightedAverageCalculatorInstance()

	var rules = wactLoadRules(t)
	var indexValue = calculator.Calculate(rules)

	assert.Equal(t, wactIndexValueExpected, indexValue, "Expecting valid index value.")
}

func wactLoadRules(t *testing.T) rule.Rules {
	var ruleProvider = rule.JSONProviderInstance()
	var allRules, rulesLoadErr = ruleProvider.GetRulesFrom("../data/rules.json")
	if nil != rulesLoadErr {
		t.Fatal("Unable to load rules:", rulesLoadErr)
	}

	var rules rule.Rules
	for ruleIdx := range allRules {
		if allRules[ruleIdx].Name == "young_dynamic" {
			rules = append(rules, allRules[ruleIdx])
		}
	}

	return rules
}
