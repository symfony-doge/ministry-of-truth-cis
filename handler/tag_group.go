// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
	"github.com/symfony-doge/ministry-of-truth-cis/response"
	"github.com/symfony-doge/ministry-of-truth-cis/tag"
)

type TagGroupGetAllResponse struct {
	response.DefaultResponse

	Payload tag.Groups `json:"tag_groups"`
}

// Tag group handler.
type tagGroupHandler struct {
	DefaultHandler

	// Provides tag groups.
	groupProvider tag.GroupProvider
}

// Returns all available tag groups.
func (h *tagGroupHandler) GetAll() gin.HandlerFunc {
	return h.handle(func(req *request.Request) interface{} {
		var payload tag.Groups = h.groupProvider.GetByLocale(req.Locale)

		return &TagGroupGetAllResponse{response.NewOkResponse(), payload}
	})
}
