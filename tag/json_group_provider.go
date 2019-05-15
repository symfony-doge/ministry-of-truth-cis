// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/symfony-doge/ministry-of-truth-cis/config"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

const (
	// Path to json file with tag groups, used by JSONGroupProvider.
	configPathDataJson string = "data.tag.group.json"
)

// Provides tag groups by json file.
type JSONGroupProvider struct {
	logger *log.Logger
}

// Returns tag groups from json file.
func (p *JSONGroupProvider) GetByLocale(locale request.Locale) Groups {
	var c = config.Instance()

	var filenameFormat = c.GetString(configPathDataJson)
	var filename = fmt.Sprintf(filenameFormat, locale)

	var buf, readErr = ioutil.ReadFile(filename)
	p.ensureEmptyError(readErr)

	var tagGroups Groups

	var unmarshalErr = json.Unmarshal(buf, &tagGroups)
	p.ensureEmptyError(unmarshalErr)

	return tagGroups
}

// Ensures error is nil or panics.
func (p *JSONGroupProvider) ensureEmptyError(err error) {
	if nil == err {
		return
	}

	if nil != p.logger {
		p.logger.Printf("JSONGroupProvider.GetByLocale: %v\n", err)
	}

	panic("tag: unable to retrieve group data.")
}

func (p *JSONGroupProvider) SetLogger(logger *log.Logger) {
	p.logger = logger
}

func NewJSONGroupProvider() *JSONGroupProvider {
	return &JSONGroupProvider{}
}
