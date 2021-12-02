package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arkrozycki/reunion/logger"
)

var log = logger.Get()

// Start function
// starts up the http server
func Start() (bool, error) {
	log.Debug().Msg("server.Start()")
	port := os.Getenv("HttpPort")

	router := Router{}
	router.New().AddRoutes()
	srv := &http.Server{
		Handler:        router.router,
		Addr:           ":" + port,
		WriteTimeout:   60 * time.Minute,
		ReadTimeout:    60 * time.Second,
		MaxHeaderBytes: 0,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err)
		}
	}()

	log.Info().Msgf("HTTP Server started %s", port)
	<-done

	return true, nil

}
