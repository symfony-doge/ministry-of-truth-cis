// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/config"
	"github.com/symfony-doge/ministry-of-truth-cis/datautil"
)

const (
	// Path to json file with rules.
	configPathRulesJson string = "data.rule.json"
)

var jsonProviderInstance *JSONProvider

var jsonProviderOnce sync.Once

// Loads rules from json file.
type JSONProvider struct {
	loader *datautil.Loader
}

func (p *JSONProvider) GetRules() (Rules, error) {
	var c = config.Instance()
	var filepath = c.GetString(configPathRulesJson)

	return p.GetRulesFrom(filepath)
}

func (p *JSONProvider) GetRulesFrom(filepath string) (Rules, error) {
	var rules Rules
	if loadErr := p.loader.LoadJSON("JSONProvider.GetRulesFrom", filepath, &rules); nil != loadErr {
		return rules, loadErr
	}

	return rules, nil
}

func NewJSONProvider() *JSONProvider {
	return &JSONProvider{
		loader: datautil.LoaderInstance(),
	}
}

func JSONProviderInstance() *JSONProvider {
	jsonProviderOnce.Do(func() {
		jsonProviderInstance = NewJSONProvider()
	})

	return jsonProviderInstance
}
