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
	// Path to json file with tags.
	configPathTagsDataJson string = "data.tag.json"
)

// Provides tags by a json file.
type JSONProvider struct {
	logger *log.Logger

	loader *datautil.Loader
}

// Returns tags from a json file.
func (p *JSONProvider) GetTags(locale request.Locale) Tags {
	var c = config.Instance()

	var filenameFormat = c.GetString(configPathTagsDataJson)
	var filename = fmt.Sprintf(filenameFormat, locale)

	var tags Tags
	if loadErr := p.loader.LoadJSON("JSONProvider.GetTags", filename, &tags); nil != loadErr {
		p.logger.Println(loadErr)

		panic("tag: unable to retrieve tags.")
	}

	return tags
}

func NewJSONProvider() *JSONProvider {
	return &JSONProvider{
		logger: DefaultLogger,
		loader: datautil.LoaderInstance(),
	}
}
