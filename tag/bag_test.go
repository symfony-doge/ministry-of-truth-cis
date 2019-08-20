// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symfony-doge/ministry-of-truth-cis/datautil"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

const (
	btTestLocale   request.Locale = "ru"
	btTestTagsJson string         = "../data/ru/tags.json"
	btTestTagName  string         = "young_dynamic"
	btTestTagTitle string         = "Динамично развивающаяся компания"
)

type btMockProvider struct {
	loader *datautil.Loader
}

func (p *btMockProvider) GetByLocale(locale request.Locale) Tags {
	var tags Tags
	var loadErr = p.loader.LoadJSON("btMockProvider.GetByLocale", btTestTagsJson, &tags)
	if nil != loadErr {
		log.Fatal(loadErr)
	}

	return tags
}

func btMockedBag() *Bag {
	var b *Bag = BagInstance()
	b.tagProvider = &btMockProvider{datautil.LoaderInstance()}

	return b
}

func TestBagGetByName(t *testing.T) {
	var b *Bag = btMockedBag()
	b.Load(btTestLocale)

	var tag, isTagExists = b.GetByName(btTestTagName, btTestLocale)

	assert.True(t, isTagExists, "Existence flag must be true for an existing tag.")
	assert.Equal(t, btTestTagTitle, tag.Title, "Expecting a valid tag data.")

	var tag2, isTagExists2 = b.GetByName("nonexistentTagName", btTestLocale)

	assert.False(t, isTagExists2, "Existence flag must be false for an non-existent tag.")
	assert.Empty(t, tag2, "Expecting a zero tag structure.")
}
