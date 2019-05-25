// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package datautil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var loaderInstance *Loader

var loaderOnce sync.Once

// Loader is a general purpose structure for boilerplate code that loads
// different resources to the memory.
type Loader struct{}

// Maps json contents to the specified variable.
func (l *Loader) LoadJSON(method, filename string, variable interface{}) error {
	var buf, readErr = ioutil.ReadFile(filename)
	if nil != readErr {
		return fmt.Errorf("%s: %v\n", method, readErr)
	}

	var unmarshalErr = json.Unmarshal(buf, &variable)
	if nil != unmarshalErr {
		return fmt.Errorf("%s: %v\n", method, unmarshalErr)
	}

	return nil
}

func LoaderInstance() *Loader {
	loaderOnce.Do(func() {
		loaderInstance = &Loader{}
	})

	return loaderInstance
}
