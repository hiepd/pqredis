package main

import (
	"pqredis/config"
	"pqredis/monitoring/logger"
)

func main() {
	l := logger.New(config.New())
	l.Infof("hello pqredis")
}
