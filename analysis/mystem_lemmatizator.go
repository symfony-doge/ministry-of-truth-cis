// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

// Performs text lemmatization via Yandex MyStem (https://tech.yandex.ru/mystem).
type MyStemLemmatizator struct{}

// Runs mystem executable with -l (lemmas only), -d (disambiguation solving)
// and -n (each word on a new line) flags.
func (l *MyStemLemmatizator) Lemmatize(input string) (output string) {
	// TODO

	return input
}
