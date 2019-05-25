// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

var bagInstance *Bag

var bagOnce sync.Once

type tagByName map[string]Tag

// Bag is a data structure that contains mappings from tag name to tag instance.
type Bag struct {
	// Loads tags from a json file.
	tagProvider *JSONProvider

	tagsByLocale map[request.Locale]tagByName
}

// Loads all available tags to the memory.
func (b *Bag) Load(locale request.Locale) {
	if _, isLocaleExists := b.tagsByLocale[locale]; isLocaleExists {
		return
	}

	var tags = b.tagProvider.GetTags(locale)
	var tbn = make(tagByName)

	for _, tag := range tags {
		tbn[tag.Name] = tag
	}

	b.tagsByLocale[locale] = tbn
}

// Returns a tag copy from the memory with given name; second retval will be
// false if a tag with specified name doesn't exists.
func (b *Bag) GetByName(name string, locale request.Locale) (Tag, bool) {
	if _, isLocaleExists := b.tagsByLocale[locale]; !isLocaleExists {
		panic("tag: bag misuse.")
	}

	var tag, isTagExists = b.tagsByLocale[locale][name]

	return tag, isTagExists
}

func BagInstance() *Bag {
	bagOnce.Do(func() {
		bagInstance = &Bag{
			tagProvider:  NewJSONProvider(),
			tagsByLocale: make(map[request.Locale]tagByName),
		}
	})

	return bagInstance
}
