// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"github.com/symfony-doge/ministry-of-truth-cis/index"
	"github.com/symfony-doge/ministry-of-truth-cis/tag"
)

var (
	defaultH  *defaultHandler
	indexH    *indexHandler
	tagGroupH *tagGroupHandler
)

func Default() *defaultHandler {
	if nil != defaultH {
		return defaultH
	}

	defaultH = &defaultHandler{}

	return defaultH
}

// Factory method for index handler.
func Index() *indexHandler {
	if nil != indexH {
		return indexH
	}

	indexH = &indexHandler{}
	indexH.defaultHandler = newDefaultHandler()

	// Index builder.
	indexH.indexBuilder = index.NewConcurrentBuilder()

	return indexH
}

// Factory method for tagGroup handler.
func TagGroup() *tagGroupHandler {
	if nil != tagGroupH {
		return tagGroupH
	}

	tagGroupH = &tagGroupHandler{}
	tagGroupH.defaultHandler = newDefaultHandler()

	// Group provider.
	var jsonGroupProvider = tag.NewJSONGroupProvider()
	jsonGroupProvider.SetLogger(tag.DefaultLogger)

	var cachedGroupProvider = tag.NewCachedGroupProvider()
	cachedGroupProvider.SetNested(jsonGroupProvider)

	tagGroupH.groupProvider = cachedGroupProvider

	return tagGroupH
}
