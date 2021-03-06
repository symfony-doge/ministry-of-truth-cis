// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"fmt"
	"log"

	"github.com/symfony-doge/ministry-of-truth-cis/config"
	"github.com/symfony-doge/ministry-of-truth-cis/datautil"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

const (
	// Path to json file with tag groups.
	configPathDataJson string = "data.tag.group.json"
)

// Provides tag groups by a json file.
type JSONGroupProvider struct {
	logger *log.Logger

	loader *datautil.Loader
}

// Returns tag groups from a json file.
func (p *JSONGroupProvider) GetByLocale(locale request.Locale) Groups {
	var c = config.Instance()

	var filenameFormat = c.GetString(configPathDataJson)
	var filename = fmt.Sprintf(filenameFormat, locale)

	var tagGroups Groups
	if loadErr := p.loader.LoadJSON("JSONGroupProvider.GetByLocale", filename, &tagGroups); nil != loadErr {
		p.logger.Println(loadErr)

		panic("tag: unable to retrieve tag groups.")
	}

	return tagGroups
}

func NewJSONGroupProvider() *JSONGroupProvider {
	return &JSONGroupProvider{
		logger: DefaultLogger,
		loader: datautil.LoaderInstance(),
	}
}
