// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

import (
	"sync"
)

// TODO: rule detransliteration.

var russianPurifierInstance *russianPurifier

var russianPurifierOnce sync.Once

// Performs word purification based on russian alphabet.
type russianPurifier struct{}

func (p *russianPurifier) Purify(word string) string {
	var validRunes []rune

	for _, r := range word {
		if 'я' >= r && r >= 'а' || 'Я' >= r && r >= 'А' {
			validRunes = append(validRunes, r)
		}
	}

	return string(validRunes)
}

func RussianPurifierInstance() *russianPurifier {
	russianPurifierOnce.Do(func() {
		russianPurifierInstance = &russianPurifier{}
	})

	return russianPurifierInstance
}
