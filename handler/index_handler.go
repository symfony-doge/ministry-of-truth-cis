// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/index"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
	"github.com/symfony-doge/ministry-of-truth-cis/response"
)

type indexRequest struct {
	request.DefaultRequest

	index.BuilderContext `json:"context" binding:"required"`
}

type indexResponse struct {
	response.DefaultResponse

	index.Index `json:"index"`
}

// Handler for index action.
type indexHandler struct {
	defaultHandler

	// Builds sanity index by specified context.
	indexBuilder index.Builder
}

// Returns all available tag groups.
func (h *indexHandler) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestFromJson indexRequest

		if err := context.ShouldBindJSON(&requestFromJson); nil != err {
			h.errorDispatcher.Dispatch(context, err)

			return
		}

		var payload *index.Index = h.indexBuilder.Build(
			requestFromJson.BuilderContext,
			requestFromJson.Locale,
		)

		context.JSON(
			http.StatusOK,
			&indexResponse{response.NewOkResponse(), *payload},
		)
	}
}
