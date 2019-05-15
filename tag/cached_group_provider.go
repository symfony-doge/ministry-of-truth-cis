// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

// Provides tag groups from memory when possible.
type CachedGroupProvider struct {
	nested GroupProvider

	cached map[request.Locale]Groups
}

func (p *CachedGroupProvider) GetByLocale(locale request.Locale) Groups {
	if fromCache, exists := p.cached[locale]; exists {
		return fromCache
	}

	var tagGroups = p.nested.GetByLocale(locale)

	p.cached[locale] = tagGroups

	return tagGroups
}

func (p *CachedGroupProvider) SetNested(nested GroupProvider) {
	p.nested = nested
}

func NewCachedGroupProvider() *CachedGroupProvider {
	return &CachedGroupProvider{cached: make(map[request.Locale]Groups)}
}
