// Package main provides ...
package main

import (
	"fmt"

	"github.com/XGHXT/dblog/pkg/config"
	"github.com/XGHXT/dblog/pkg/core/backup/ping"
	"github.com/XGHXT/dblog/pkg/core/backup/timer"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi, it's App " + config.Conf.BackupApp.Name)

	endRun := make(chan error, 1)

	runTimer(endRun)

	runHTTPServer(endRun)
	fmt.Println(<-endRun)
}

func runTimer(endRun chan error) {
	go func() {
		endRun <- timer.Start()
	}()
}

func runHTTPServer(endRun chan error) {
	if !config.Conf.BackupApp.EnableHTTP {
		return
	}

	if config.Conf.RunMode == config.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()

	// route
	ping.RegisterRoutes(e)

	// start
	address := fmt.Sprintf(":%d", config.Conf.BackupApp.HTTPPort)
	go func() {
		endRun <- e.Run(address)
	}()
	fmt.Println("HTTP server running on: " + address)
}
