// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"sync"

	"github.com/kljensen/snowball/russian"
)

var snowballStemmerInstance *snowballStemmer

var snowballStemmerOnce sync.Once

// Performs word stemming.
type snowballStemmer struct{}

func (l *snowballStemmer) Stem(word string) (string, error) {
	var wordStemmed = russian.Stem(word, true)

	// TODO: remove unwanted symbols, then fix test [0].

	return wordStemmed, nil
}

func SnowballStemmerInstance() *snowballStemmer {
	snowballStemmerOnce.Do(func() {
		snowballStemmerInstance = &snowballStemmer{}
	})

	return snowballStemmerInstance
}
