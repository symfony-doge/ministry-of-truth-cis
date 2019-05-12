// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"strconv"

	"github.com/symfony-doge/ministry-of-truth-cis/handler/tag/group"

	"github.com/gin-gonic/gin"
)

var (
	GinMode    *string = flag.String("mode", gin.DebugMode, "Gin mode: debug, test or release.")
	ServerPort *int    = flag.Int("port", 9595, "Port to listen.")
)

func init() {
	flag.Parse()

	gin.SetMode(*GinMode)
}

func main() {
	var router = gin.Default()

	router.GET("/tag/groups", handler.TagGroupGetAll)

	var addr = ":" + strconv.Itoa(*ServerPort)
	router.Run(addr)
}
