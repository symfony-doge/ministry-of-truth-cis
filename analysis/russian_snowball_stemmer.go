// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"sync"

	"github.com/kljensen/snowball/russian"
)

var russianSnowballStemmerInstance *russianSnowballStemmer

var russianSnowballStemmerOnce sync.Once

// Performs word stemming.
type russianSnowballStemmer struct{}

func (l *russianSnowballStemmer) Stem(word string) (string, error) {
	var wordStemmed = russian.Stem(word, true)

	return wordStemmed, nil
}

func RussianSnowballStemmerInstance() *russianSnowballStemmer {
	russianSnowballStemmerOnce.Do(func() {
		russianSnowballStemmerInstance = &russianSnowballStemmer{}
	})

	return russianSnowballStemmerInstance
}
