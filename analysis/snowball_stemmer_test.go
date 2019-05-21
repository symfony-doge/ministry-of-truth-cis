// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sstTestString = "Молодая, динамично РаЗВИваюЩАяся КОМПАНИЯ"
)

var (
	sstSnowballStemmer *snowballStemmer = SnowballStemmerInstance()
	sstResultsExpected                  = [...]string{
		"молодая,",
		"динамичн",
		"развива",
		"компан",
	}
)

func TestStemmerStemSnowball(t *testing.T) {
	var testStrings = strings.Fields(sstTestString)

	for idx := range testStrings {
		var testResult, err = sstSnowballStemmer.Stem(testStrings[idx])

		assert.NoError(t, err, "Expecting no error.")
		assert.Equal(t, sstResultsExpected[idx], testResult, "Expecting a valid stem.")
	}
}

func BenchmarkStemmerStemSnowball(b *testing.B) {
	var testStrings = strings.Fields(sstTestString)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		for idx := range testStrings {
			sstSnowballStemmer.Stem(testStrings[idx])
		}
	}
}

// $ go test ./analysis -bench StemSnowball -benchmem
// 21794 ns/op    4631 B/op    384 allocs/op
