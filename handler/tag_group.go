// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
	"github.com/symfony-doge/ministry-of-truth-cis/response"
	"github.com/symfony-doge/ministry-of-truth-cis/tag"
)

type TagGroupsGetAllResponse struct {
	response.DefaultResponse

	Payload tag.Groups `json:"tag_groups"`
}

// Tag group handler.
type tagGroupHandler struct {
	// Converts body of HTTP request into appropriate request.Request structure.
	requestBinder request.Binder

	// Provides tag groups.
	groupProvider tag.GroupProvider
}

// Returns all available tag groups.
func (handler *tagGroupHandler) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req *request.Request = handler.requestBinder.Bind(context)
		var payload tag.Groups = handler.groupProvider.GetByLocale(req.Locale)

		var response = &TagGroupsGetAllResponse{response.NewOkResponse(), payload}

		context.JSON(http.StatusOK, response)
	}
}
