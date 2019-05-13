// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package tag

type Groups []Group

// Represents group of tags.
type Group struct {
	Name        string `json: "name"`
	Description string `json: "description"`
	Color       string `json: "color"`
}
