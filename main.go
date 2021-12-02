package main

import (
	"github.com/arkrozycki/reunion/datastore"
	logger "github.com/arkrozycki/reunion/logger"
	"github.com/arkrozycki/reunion/server"
)

var log = logger.Get()
var db = datastore.GetDatastore()

// main function
func main() {
	var err error
	//conf.Config()
	log.Info().Msg("reunion starting ...")

	exit, err := server.Start()
	if err != nil {
		log.Fatal().Err(err)
	}

	if exit {
		// start graceful shutdown
		log.Info().Msg("Service graceful shutdown started ...")
	}

}
