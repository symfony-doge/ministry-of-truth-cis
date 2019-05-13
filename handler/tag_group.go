// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tagGroup struct{}

// Returns all available tag groups.
func (handler *tagGroup) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"status": "OK",
				"errors": []string{},
				"tag_groups": []gin.H{
					{
						"name":        "soft",
						"description": "A senseless verbiage, poor language or just a spam words",
						"color":       "#61C3FF",
					},
				},
			},
		)
	}
}

func TagGroup() *tagGroup {
	return &tagGroup{}
}
