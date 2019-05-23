// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

// Provides tag groups from memory when possible; thread-safe.
type CachedGroupProvider struct {
	// Actual data provider if cache get is missed.
	nested GroupProvider

	// For concurrent execution safety.
	// Guarantees that only first cache set action will take effect
	// to all goroutines where this instance is referenced.
	mu sync.Mutex
	// In-memory storage, protected by mu.
	cached map[request.Locale]Groups
}

func (p *CachedGroupProvider) GetByLocale(locale request.Locale) Groups {
	if fromCache, exists := p.cached[locale]; exists {
		return fromCache
	}

	return p.warmup(locale)
}

func (p *CachedGroupProvider) warmup(locale request.Locale) Groups {
	p.mu.Lock()
	defer p.mu.Unlock()

	if existentValue, isAlreadySet := p.cached[locale]; isAlreadySet {
		return existentValue
	}

	p.cached[locale] = p.nested.GetByLocale(locale)

	return p.cached[locale]
}

func (p *CachedGroupProvider) SetNested(nested GroupProvider) {
	p.nested = nested
}

func NewCachedGroupProvider() *CachedGroupProvider {
	return &CachedGroupProvider{
		cached: make(map[request.Locale]Groups),
		mu:     sync.Mutex{},
	}
}
