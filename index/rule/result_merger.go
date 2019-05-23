// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

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
		result.eventContextByWord = make(map[string]OccurrenceFoundContext)
	}

	// We don't need to perform any actions if rule is already marked
	// as applicable to the text.
	if result.isMergeConditionsMet {
		return
	}

	// We don't need to process the same word twice if such occurs.
	if _, isContextForWordExists := result.eventContextByWord[context.word]; isContextForWordExists {
		return
	}

	result.eventContextByWord[context.word] = context

	if m.checkMergeConditions(&result) {
		result.isMergeConditionsMet = true
	}

	m.resultByRuleName[rule.Name] = result
}

// Returns positive whenever all merge conditions
// from rule specification are met.
func (m *MatchTaskResultMerger) checkMergeConditions(result *matchTaskResult) bool {
	// TODO check rule specification.
	// len fast retval cond.

	return false
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
