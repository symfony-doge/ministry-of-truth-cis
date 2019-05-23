// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
	"os"
)

// Package-level logger.
var DefaultLogger *log.Logger = log.New(os.Stdout, "[rule] ", log.Ldate|log.Ltime|log.Lshortfile)

// MergeCondition is a criteria for merging positive results of words
// occurrence checks.
type MergeCondition struct {
	// 0..N value means maximum allowed offset between target word and previous word
	// in specification for "merging" positive results of occurrence check.
	// Negative value means no offset checks are required.
	OffsetPreviousMax int `json:"offsetPreviousMax"`
}

// Describes words to be checked for occurrence in the text sentence.
type SpecificationEntry struct {
	// One of the words will be enough for positive occurrence check result.
	Words []string `json:"words"`

	// A set of context markers; one is enough for positive occurrence check result.
	Contexts []string `json:"contexts"`

	// Additional criteria for "merging" positive occurrence checks.
	MergeConditions []MergeCondition `json:"conditions"`
}

type Rules []*Rule

// func (r Rules) String() string {
// 	var str strings.Builder
// 	str.WriteString(fmt.Sprintf("%v\n", r[0].Weight))
// 	return str.String()
// }

// Rule for sanity index calculation. Rule is applicable to the text sentence if:
// 1. One of the specified rule contexts equals to "contextMarker" for the text sentence.
// 2. A text sentence contains all words in the rule specification.
// 3. All merge conditions for matched words are met (e.g. some words should
// not have any symbols between).
type Rule struct {
	// Rule name.
	Name string `json:"name"`

	// Rule "cost" for sanity index calculation algorithm.
	Weight float64 `json:"weight"`

	// Rule specification is a set of configuration entries that describes
	// cases when this rule is applicable to the text sentence.
	Specification []SpecificationEntry `json:"specification"`

	TagNames []string `json:"tags"`
}
