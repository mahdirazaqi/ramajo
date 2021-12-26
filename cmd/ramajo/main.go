package main

import (
	"flag"
	"log"

	"github.com/mahdirazaqi/ramajo/config"
)

var configPath = flag.String("c", "./config.json", "config file path")

func main() {
	flag.Parse()

	_, err := config.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}
}
