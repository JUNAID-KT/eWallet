package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JUNAID-KT/eWallet/app"
	routing "github.com/JUNAID-KT/eWallet/router"
	"github.com/JUNAID-KT/eWallet/search_engine"
	log "github.com/Sirupsen/logrus"
)

func main() {

	app.Init()
	elasticClient := search_engine.GetESInstance()
	if elasticClient == nil {
		log.Fatalln("Connection failure.Unable to connect to Elasticsearch.")
	} else {
		defer elasticClient.Stop()
	}
	router := routing.InitRoutes()
	// Create the Server
	server := &http.Server{
		Addr:    app.Config.Server,
		Handler: router,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Running the HTTP server

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithFields(log.Fields{"error": err.Error()}).Fatalf("server listen failed")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.WithFields(log.Fields{"method": "main"}).Infoln("server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Fatalf("server shutdown failed")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.WithFields(log.Fields{"method": "main"}).Infoln("server exited")
}
