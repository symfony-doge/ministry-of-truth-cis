// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/config"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

const (
	// Path to json file with tag groups, used by JSONGroupProvider.
	configurationParameterDataJSON = "data.tag.group.json"
)

type GroupProvider interface {
	// GetByLocale panics if it cannot retreive data.
	GetByLocale(request.Locale) Groups
}

// Provides tag groups by json file.
type JSONGroupProvider struct {
	logger *log.Logger

	// TODO cache
	cached map[request.Locale]Groups
}

// Returns tag groups stored in json file.
func (p *JSONGroupProvider) GetByLocale(locale request.Locale) Groups {
	var c = config.Instance()

	var filenameFormat = c.GetString(configurationParameterDataJSON)
	var filename = fmt.Sprintf(filenameFormat, locale)
	var jsonFile, openErr = os.Open(filename)
	p.ensureEmptyError(openErr)

	defer jsonFile.Close()

	var buf = &bytes.Buffer{}
	_, readErr := io.Copy(buf, jsonFile)
	p.ensureEmptyError(readErr)

	var tagGroups Groups
	var unmarshalErr = json.Unmarshal(buf.Bytes(), &tagGroups)
	p.ensureEmptyError(unmarshalErr)

	if gin.IsDebugging() && nil != p.logger {
		p.logger.Printf("JSONGroupProvider.GetByLocale: %v\n", tagGroups)
	}

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

	panic("tag: unable to retreive group data.")
}

func (p *JSONGroupProvider) SetLogger(logger *log.Logger) {
	p.logger = logger
}
