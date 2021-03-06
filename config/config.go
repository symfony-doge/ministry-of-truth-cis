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

	jjw "github.com/spf13/jwalterweatherman"
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

func init() {
	// Category for logger that is used by viper.
	jjw.SetPrefix("Viper")
}

// Loads configuration according to specified Gin mode.
func Load(mode string) error {
	return LoadFrom(mode, []string{"config"})
}

func LoadFrom(mode string, configPaths []string) error {
	config = viper.New()

	for cpIdx := range configPaths {
		config.AddConfigPath(configPaths[cpIdx])
	}

	config.SetConfigName(mode)
	config.SetConfigType("yaml")

	config.AutomaticEnv()

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
