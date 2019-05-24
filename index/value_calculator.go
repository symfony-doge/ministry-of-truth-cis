// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
)

// valueCalculator is an interface for components that provide algorithms
// for index value calculation, based on rules, determined by the input analysis.
type valueCalculator interface {
	Calculate(rule.Rules) float64
}
