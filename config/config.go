// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package config implements application configuration storage.
// Parameters are loaded according to specified running mode
// (-mode can be debug, test or release).
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// ConfigNotFoundError is returned by Load if it fails.
type ConfigNotLoadedError struct {
	mode string
}

// ConfigNotLoadedError.Error returns the formatted error.
func (err ConfigNotLoadedError) Error() string {
	return fmt.Sprintf("Unable to load configuration for mode %q.", err.mode)
}

// Loads configuration according to specified Gin mode.
func Load(mode string) error {
	config = viper.New()

	config.AddConfigPath("config")
	config.SetConfigName(mode)
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); nil != err {
		log.Println(err)

		return ConfigNotLoadedError{mode}
	}

	return nil
}

// Returns config instance.
func Instance() *viper.Viper {
	return config
}
