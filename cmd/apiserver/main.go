package main

import (
	"flag"
	"log"

	"github.com/exdev-studio/requests-dashboard-api/internal/apiserver"
)

var (
	bindAddr string
	logLevel string
)

func init() {
	flag.StringVar(&bindAddr, "bind-addr", ":8080", "bind address")
	flag.StringVar(&logLevel, "log-level", "info", "log level")
}

func main() {
	flag.Parse()

	c := apiserver.NewConfig()
	if bindAddr != "" {
		c.BindAddr = bindAddr
	}

	if logLevel != "" {
		c.LogLevel = logLevel
	}

	if err := apiserver.Start(c); err != nil {
		log.Fatal(err)
	}
}
