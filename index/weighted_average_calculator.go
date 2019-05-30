// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
)

var weightedAverageCalculatorI *weightedAverageCalculator

var weightedAverageCalculatorOnce sync.Once

// weightedAverageCalculator is a simple value calculator that provides
// sanity index as a percentage of the sum of grades (g) from all applicable
// rules multiplied by their weights (w) and divided by the weights sum (P).
// (g1*w1 + g2*w2 + ... + gN*wN) / P; where N is a rule number.
type weightedAverageCalculator struct{}

// Implements valueCalculator interface.
func (vc *weightedAverageCalculator) Calculate(rules rule.Rules) float64 {
	if len(rules) < 1 {
		return 99.9
	}

	var weightedGradeSum, weightTotal float64 = 0.0, 0.0

	for ruleIdx := range rules {
		weightedGradeSum += rules[ruleIdx].Grade * rules[ruleIdx].Weight
		weightTotal += rules[ruleIdx].Weight
	}

	var value = weightedGradeSum / weightTotal

	if value > 99.9 {
		value = 99.9
	}

	return value
}

func newWeightedAverageCalculator() *weightedAverageCalculator {
	return &weightedAverageCalculator{}
}

func weightedAverageCalculatorInstance() *weightedAverageCalculator {
	weightedAverageCalculatorOnce.Do(func() {
		weightedAverageCalculatorI = newWeightedAverageCalculator()
	})

	return weightedAverageCalculatorI
}
