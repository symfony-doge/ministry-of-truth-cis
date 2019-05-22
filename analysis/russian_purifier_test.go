// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	rptTestString = "Мол:одая, дин;амично.. РаЗВИваюЩАяс!?я КО~МПАНИЯ"
)

var (
	rptRussianPurifier *russianPurifier = RussianPurifierInstance()
	rptResultsExpected                  = [...]string{
		"Молодая",
		"динамично",
		"РаЗВИваюЩАяся",
		"КОМПАНИЯ",
	}
)

func TestPurifierStemRussianSnowball(t *testing.T) {
	var testStrings = strings.Fields(rptTestString)

	for idx := range testStrings {
		var testResult = rptRussianPurifier.Purify(testStrings[idx])

		assert.Equal(t, rptResultsExpected[idx], testResult, "Expecting a purified word.")
	}
}
