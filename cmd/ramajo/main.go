package main

import (
	"flag"

	"github.com/mahdirazaqi/ramajo/config"
	"github.com/mahdirazaqi/ramajo/internal/server"
	"github.com/mahdirazaqi/ramajo/pkg/logger"
)

var configPath = flag.String("c", "./config.json", "config file path")

func main() {
	flag.Parse()

	c, err := config.Load(*configPath)
	if err != nil {
		logger.Error(err)
	}

	if err := server.Start(c.Port); err != nil {
		logger.Error(err)
	}
}
