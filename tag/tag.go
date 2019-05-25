// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

import (
	"log"
	"os"
)

// Package-level logger.
var DefaultLogger *log.Logger = log.New(os.Stdout, "[tag] ", log.Ldate|log.Ltime|log.Lshortfile)

type Tags []Tag

// Represents search tag.
type Tag struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	ImageUrl    string `json:"image_url"`
	GroupName   string `json:"group"`
}
