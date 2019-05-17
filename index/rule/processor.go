// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// Determines which rules has match for a given text
// and returns a result set.
type Processor interface {
	FindMatch(string) (Rules, error)
}

// Can be thrown by a processor that uses a rule events system;
// indicates that an event listener has not been initialized.
type EventListenerNotStartedError struct{}

func (err EventListenerNotStartedError) Error() string {
	return "Unable to FindMatch. Event listener is not started."
}
