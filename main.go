package main

import (
	"pqredis/config"
	"pqredis/monitoring/logger"
	"pqredis/postgres"
)

func main() {
	l := logger.New(config.New())
	l.Infof("hello pqredis")

	cfg := config.New()
	_, err := postgres.New(cfg, l)
	if err != nil {
		l.Errorf(err.Error())
	} else{
		l.Infof("DB is set up")
	}
}
