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

const (
	configPathLogMain  string = "log.main.filepath"
	configPathLogError string = "log.error.filepath"
)

// LoggerNotCreatedError is returned if NewErrorLogger call failed.
type LoggerNotCreatedError struct {
	prefix string
}

// LoggerNotCreatedError implements error interface.
func (err LoggerNotCreatedError) Error() string {
	return fmt.Sprintf("Unable to create logger (prefix='%s').", err.prefix)
}

// WriterNotCreatedError is returned if NewWriter or NewErrorWriter call failed.
type WriterNotCreatedError struct {
	filepath string
}

// WriterNotCreatedError implements error interface.
func (err WriterNotCreatedError) Error() string {
	return fmt.Sprintf("Unable to create log writer (filename='%s').", err.filepath)
}

// Creates new logger for error messages.
func NewErrorLogger(prefix string) (*golog.Logger, error) {
	var writer, err = NewErrorWriter()
	if nil != err {
		golog.Println(err)

		return nil, LoggerNotCreatedError{prefix}
	}

	var logger = golog.New(*writer, prefix, golog.Ldate|golog.Ltime|golog.Lshortfile)

	return logger, nil
}

// Creates new stream writer dedicated to write messages.
func NewWriter() (*io.Writer, error) {
	var filepath = resolveFilepath(configPathLogMain)

	var file, err = os.Create(filepath)
	if nil != err {
		golog.Println(err)

		return nil, WriterNotCreatedError{filepath}
	}

	var writer = io.MultiWriter(file, os.Stdout)

	return &writer, nil
}

// Creates new stream writer dedicated to write error messages.
func NewErrorWriter() (*io.Writer, error) {
	var filepath = resolveFilepath(configPathLogError)

	var file, err = os.Create(filepath)
	if nil != err {
		golog.Println(err)

		return nil, WriterNotCreatedError{filepath}
	}

	var writer = io.MultiWriter(file, os.Stderr)

	return &writer, nil
}

// Returns filename for writer.
func resolveFilepath(configPath string) (filepath string) {
	var c = config.Instance()
	var filepathFormat = c.GetString(configPath)

	filepath = fmt.Sprintf(filepathFormat, gin.Mode())

	return
}
