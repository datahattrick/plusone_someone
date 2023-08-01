package main

import (
	"log"

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

	// LDAP Scrape
	err = utils.LDAPStartTLS(&cfg)
	if err != nil {
		log.Println("Failed to connect to LDAP server and gather users", err)
	}

	//startServer
	http.StartServer(&cfg)
}
