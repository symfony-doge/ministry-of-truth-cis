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
	"github.com/symfony-doge/ministry-of-truth-cis/index/rule"
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
		log.Fatal("Unable to load config:", err)
	}

	// Setting up log files.
	var logWriter, lwErr = applog.NewWriter()
	if nil != lwErr {
		log.Fatal("Unable to init main log:", lwErr)
	}
	gin.DefaultWriter = *logWriter

	var errLogWriter, elwErr = applog.NewErrorWriter()
	if nil != elwErr {
		log.Fatal("Unable to init error log:", elwErr)
	}
	gin.DefaultErrorWriter = *errLogWriter

	// Warming up rules index for matching word occurrences (index action).
	if riBuildErr := rule.InvertedIndexInstance().Build(); nil != riBuildErr {
		log.Fatal("Unable to build rule index: ", riBuildErr)
	}
}

func configureRouter() *gin.Engine {
	var router *gin.Engine = gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Recovery())

	router.NoRoute(handler.Default().RouteNotFound())
	router.NoMethod(handler.Default().MethodNotAllowed())

	router.GET("/tag/groups", handler.TagGroup().GetAll())
	router.POST("/tag/groups", handler.TagGroup().GetAll())
	router.POST("/index", handler.Index().Index())

	return router
}

func main() {
	var router = configureRouter()

	var networkAddressToListen = ":" + strconv.Itoa(ServerPort)
	var err = router.Run(networkAddressToListen)

	log.Fatal(err)
}
