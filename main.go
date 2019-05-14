// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	jjw "github.com/spf13/jwalterweatherman"
	"github.com/symfony-doge/ministry-of-truth-cis/config"
	"github.com/symfony-doge/ministry-of-truth-cis/handler"
	applog "github.com/symfony-doge/ministry-of-truth-cis/log"
	"github.com/symfony-doge/ministry-of-truth-cis/middleware"
)

var (
	GinMode    string
	ServerPort int
)

func init() {
	// Processing command-line arguments.
	flag.StringVar(&GinMode, "mode", gin.DebugMode, "Gin mode: debug, test or release.")
	flag.IntVar(&ServerPort, "port", 9595, "Port to listen.")
	flag.Parse()

	gin.SetMode(GinMode)

	// Adjusting log level for Viper configuration manager in debug mode.
	if gin.IsDebugging() {
		jjw.SetLogThreshold(jjw.LevelTrace)
		jjw.SetStdoutThreshold(jjw.LevelTrace)
	}

	// Loading application configuration.
	if err := config.Load(GinMode); nil != err {
		log.Fatal(err)
	}

	// Setting up log files.
	var logWriter, lwErr = applog.NewWriter()
	if nil != lwErr {
		log.Println(lwErr)
		log.Fatal("Unable to init main log.")
	}
	gin.DefaultWriter = *logWriter

	var errLogWriter, elwErr = applog.NewErrorWriter()
	if nil != elwErr {
		log.Println(elwErr)
		log.Fatal("Unable to init error log.")
	}
	gin.DefaultErrorWriter = *errLogWriter
}

func main() {
	var router *gin.Engine = gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Recovery())

	var tagRouterGroup *gin.RouterGroup = router.Group("/tag")
	{
		var tagGroupRouterGroup *gin.RouterGroup = tagRouterGroup.Group("/groups")
		{
			tagGroupRouterGroup.GET("", handler.TagGroup().GetAll())
			tagGroupRouterGroup.POST("", handler.TagGroup().GetAll())
		}
	}

	var networkAddressToListen = ":" + strconv.Itoa(ServerPort)
	var err = router.Run(networkAddressToListen)

	log.Fatal(err)
}
