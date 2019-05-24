// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/symfony-doge/ministry-of-truth-cis/config"
)

const (
	// Path to json file with rules.
	configPathRulesJson string = "data.rule.json"
)

var jsonProviderInstance *JSONProvider

var jsonProviderOnce sync.Once

// Loads rules from json file.
type JSONProvider struct{}

func (l *JSONProvider) GetRules() (Rules, error) {
	var c = config.Instance()
	var filepath = c.GetString(configPathRulesJson)

	return l.GetRulesFrom(filepath)
}

func (l *JSONProvider) GetRulesFrom(filepath string) (Rules, error) {
	var buf, readErr = ioutil.ReadFile(filepath)
	if nil != readErr {
		return nil, readErr
	}

	var rules Rules
	if unmarshalErr := json.Unmarshal(buf, &rules); nil != unmarshalErr {
		return nil, unmarshalErr
	}

	return rules, nil
}

func NewJSONProvider() *JSONProvider {
	return &JSONProvider{}
}

func JSONProviderInstance() *JSONProvider {
	jsonProviderOnce.Do(func() {
		jsonProviderInstance = NewJSONProvider()
	})

	return jsonProviderInstance
}
