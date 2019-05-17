// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"sync"
)

var myStemLemmatizerInstance *myStemLemmatizer

var myStemLemmatizerOnce sync.Once

// Performs text lemmatization via Yandex MyStem (https://tech.yandex.ru/mystem).
type myStemLemmatizer struct{}

// Runs mystem executable with -l (lemmas only), -d (disambiguation solving)
// and -n (each word on a new line) flags.
func (l *myStemLemmatizer) Lemmatize(input string) (string, error) {
	// TODO

	return input, nil
}

func MyStemLemmatizerInstance() *myStemLemmatizer {
	myStemLemmatizerOnce.Do(func() {
		myStemLemmatizerInstance = &myStemLemmatizer{}
	})

	return myStemLemmatizerInstance
}
