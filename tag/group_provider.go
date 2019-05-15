// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

// Provides tag groups.
type GroupProvider interface {
	// GetByLocale panics if it cannot retrieve data.
	GetByLocale(request.Locale) Groups
}
