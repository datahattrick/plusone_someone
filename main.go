package main

import (
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

	//Read config this will generate defaults if fail
	utils.LoadDotEnvFile()
	// Connect to the database
	utils.ConnectDB()

	//startServer
	http.StartServer()
}
