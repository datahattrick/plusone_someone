package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/datahattrick/plusone_someone/docs"
	"github.com/datahattrick/plusone_someone/internal/http"
	"github.com/datahattrick/plusone_someone/internal/logger"
	"github.com/datahattrick/plusone_someone/internal/utils"
)

// @title			Plusone Someone API
// @version		0.1
// @description	A simple API to create a message and give someone a plusone.
//
// @schemes		http https
// @host			localhost:8000
// @BasePath		/v1
func main() {
	l := logger.New("goapp", "v1.0.0", 1)
	logger.UpdateDefaultLogger(l)

	cfg := utils.Config{}

	//Read config this will generate defaults if fail
	err := utils.LoadDotEnvFile(&cfg)
	if err != nil {
		log.Fatal("Need settings", err)
	}
	// Connect to the database
	utils.ConnectDB()

	// LDAP Scrape don't wait for it
	go func() {
		err = utils.LDAPStartTLS(&cfg)
		if err != nil {
			log.Println("Failed to connect to LDAP server and gather users", err)
		}
	}()

	server := http.NewServerHTTP(&cfg)

	// Channel for sending server shutdown signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Start server in seperate go routine
	go func() {
		server.Start(&cfg)
	}()

	// recieving server shutdown signal
	sig := <-signalCh
	log.Println("Received signal : ", sig)

	// Graceful shutdown
	if err := server.ShutDown(); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

}
