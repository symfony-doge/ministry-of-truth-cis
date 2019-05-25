// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

var aggregatorInstance *Aggregator

var aggregatorOnce sync.Once

// Performs tags aggregation ops.
type Aggregator struct {
	// A bag with tags :)
	tagBag *Bag
}

func (a *Aggregator) ExtractTagNames(rules rule.Rules) []string {
	var tagNames = []string{}
	var tagMap = make(map[string]bool)

	for _, rule := range rules {
		for _, tagName := range rule.TagIdentifiers {
			if _, isTagNameExists := tagMap[tagName]; !isTagNameExists {
				tagNames = append(tagNames, tagName)
				tagMap[tagName] = true
			}
		}
	}

	return tagNames
}

func (a *Aggregator) AggregateByGroup(tagNames []string, locale request.Locale) map[string]Tags {
	a.tagBag.Load(locale)

	var aggregated = make(map[string]Tags)

	for _, tagName := range tagNames {
		var tag, isTagExists = a.tagBag.GetByName(tagName, locale)
		if !isTagExists {
			panic("tag: invalid tag name.")
		}

		if _, isGroupNameExists := aggregated[tag.GroupName]; isGroupNameExists {
			aggregated[tag.GroupName] = append(aggregated[tag.GroupName], tag)
		} else {
			aggregated[tag.GroupName] = Tags{tag}
		}
	}

	return aggregated
}

func AggregatorInstance() *Aggregator {
	aggregatorOnce.Do(func() {
		aggregatorInstance = &Aggregator{
			tagBag: BagInstance(),
		}
	})

	return aggregatorInstance
}
