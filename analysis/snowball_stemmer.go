// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"strings"
	"sync"

	"github.com/kljensen/snowball/russian"
)

var snowballStemmerInstance *snowballStemmer

var snowballStemmerOnce sync.Once

// Performs text stemming.
type snowballStemmer struct{}

func (l *snowballStemmer) Stem(input string) ([]string, error) {
	var fields = strings.Fields(input)
	var outputStemmed []string

	for idx := range fields {
		outputStemmed = append(outputStemmed, russian.Stem(fields[idx], true))
	}

	return outputStemmed, nil
}

func SnowballStemmerInstance() *snowballStemmer {
	snowballStemmerOnce.Do(func() {
		snowballStemmerInstance = &snowballStemmer{}
	})

	return snowballStemmerInstance
}
