// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package log

import (
	"fmt"
	"io"
	golog "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/config"
)

// LogNotConfiguredError is returned if either Configurator.ConfigureCategoryMain
// or Configurator.ConfigureCategoryError fails.
type LogNotConfiguredError struct {
	category string
}

func (err LogNotConfiguredError) Error() string {
	return fmt.Sprintf("Unable to configure %q log file.", err.category)
}

type Configurator struct{}

// Configures logging for all available categories.
func (configurator *Configurator) ConfigureAllCategories() error {
	if err := configurator.ConfigureCategoryMain(); nil != err {
		return err
	}

	return configurator.ConfigureCategoryError()
}

// Configures message logging for main category.
func (configurator *Configurator) ConfigureCategoryMain() error {
	return configurator.ConfigureCategory(
		"main",
		func(file *os.File) {
			gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		},
	)
}

// Configures message logging for error category.
func (configurator *Configurator) ConfigureCategoryError() error {
	return configurator.ConfigureCategory(
		"error",
		func(file *os.File) {
			gin.DefaultErrorWriter = io.MultiWriter(file, os.Stdout)
		},
	)
}

// Configures message logging for specified category.
func (configurator *Configurator) ConfigureCategory(
	category string,
	callback func(*os.File),
) error {
	var c = config.Instance()

	var logPathFormatParameter = fmt.Sprintf("log.%s.filepath", category)
	var logPathFormat = c.GetString(logPathFormatParameter)
	var logPath = fmt.Sprintf(logPathFormat, gin.Mode())

	var logFile, err = os.Create(logPath)

	if nil != err {
		golog.Println(err)

		return LogNotConfiguredError{category}
	}

	callback(logFile)

	return nil
}
