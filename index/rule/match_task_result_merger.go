// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"math"
)

// NOTE: Current implementation doesn't support same words under multiple context markers (todo).
// NOTE: Offset condition is unstable; word dublicates is not supported yet (todo).

// Represents a partial result of task execution.
type matchTaskResult struct {
	// Reference to the rule.
	rule *Rule

	// Words found by workers mapped to the related event context.
	eventContextByWord map[string]OccurrenceFoundContext

	// Becomes positive and indicates that rule is applicable to the text if:
	// 1) All words declared by rule specification are present.
	// 2) All merge conditions for specification entries are met
	// (e.g. max offset between words is not exceeded).
	isMergeConditionsMet bool
}

type resultByRuleName map[string]matchTaskResult

// ResultMerger prepares a total result of task execution in concurrent
// environment; result parts will be received and "merged" from the workers.
type MatchTaskResultMerger struct {
	// Maps rule name to the task execution result.
	resultByRuleName
}

func (m *MatchTaskResultMerger) Merge(rule *Rule, context OccurrenceFoundContext) {
	var result, isResultExists = m.resultByRuleName[rule.Name]
	if !isResultExists {
		result.rule = rule
		result.eventContextByWord = make(map[string]OccurrenceFoundContext)
	}

	// We don't need to perform any actions if rule is already marked
	// as applicable to the text.
	if result.isMergeConditionsMet {
		return
	}

	// We don't need to process the same word twice.
	if c, isContextForWordExists := result.eventContextByWord[context.wordProcessed]; isContextForWordExists {
		// Unless an occurrence case with lower position will found.
		if c.offset < context.offset {
			return
		}
	}

	result.eventContextByWord[context.wordProcessed] = context

	if m.checkMergeConditions(result) {
		result.isMergeConditionsMet = true
	}

	m.resultByRuleName[rule.Name] = result
}

// Returns positive whenever all merge conditions
// from rule specification are met.
func (m *MatchTaskResultMerger) checkMergeConditions(result matchTaskResult) bool {
	// Occurrence for all words from specification is required.
	if len(result.eventContextByWord) < len(result.rule.Specification) {
		return false
	}

	var prevSpecEntry *SpecificationEntry

	for specEntryIdx := range result.rule.Specification {
		if !m.between(result, prevSpecEntry, &result.rule.Specification[specEntryIdx]) {
			return false
		}

		prevSpecEntry = &result.rule.Specification[specEntryIdx]
	}

	return true
}

// Check merge conditions between two concrete partial results.
func (m *MatchTaskResultMerger) between(
	result matchTaskResult,
	prev *SpecificationEntry,
	cur *SpecificationEntry,
) bool {
	if nil == prev {
		return true
	}

	var prevResultContext = m.resolveResultContextForSpec(&result, prev)
	var curResultContext = m.resolveResultContextForSpec(&result, cur)

	// Checking offset between words.
	if cur.MergeConditions.OffsetPreviousMax > 0 {
		var offsetBetweenWords = int(math.Abs(float64(curResultContext.offset - prevResultContext.offset)))

		if offsetBetweenWords > cur.MergeConditions.OffsetPreviousMax {
			return false
		}
	}

	return true
}

func (m *MatchTaskResultMerger) resolveResultContextForSpec(
	result *matchTaskResult,
	specificationEntry *SpecificationEntry,
) OccurrenceFoundContext {
	for _, word := range specificationEntry.Words {
		if resultContext, isContextExists := result.eventContextByWord[word]; isContextExists {
			return resultContext
		}
	}

	panic("rule: match task results are out of sync with rule spec.")
}

func (m *MatchTaskResultMerger) GetResult() Rules {
	var rules Rules

	for _, result := range m.resultByRuleName {
		if result.isMergeConditionsMet {
			rules = append(rules, result.rule)
		}
	}

	return rules
}

func NewMatchTaskResultMerger() *MatchTaskResultMerger {
	return &MatchTaskResultMerger{
		resultByRuleName: make(resultByRuleName),
	}
}
