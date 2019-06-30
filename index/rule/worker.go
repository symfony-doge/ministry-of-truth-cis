// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"context"

	"github.com/symfony-doge/event"
)

// Performs rule processing routine.
type Worker interface {
	// Sets context of partial task for processing.
	SetContext(context.Context)

	// Adds a channel or group of channels for worker's events.
	AddNotifyChannel(...chan<- event.Event)

	// Starts routine execution.
	Run()
}
