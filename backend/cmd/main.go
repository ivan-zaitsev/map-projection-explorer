package main

import (
	"log"
	"map-projection-explorer-backend/internal/repository"
	http_server "map-projection-explorer-backend/internal/server/http"
	"map-projection-explorer-backend/internal/service"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	epsgRepository repository.EpsgExtentRepository
	srsRepository  repository.SrsRepository
	crsService     service.CrsService
)

func main() {
	setupConfig()
	setupRepositories()
	setupServices()
	setupHttpServer()
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.BindEnv("database.uri", "DATABASE_URI")
}

func setupRepositories() {
	db, err := repository.NewDatabase(viper.GetString("database.uri"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	epsgRepository = repository.NewEpsgExtentRepository(db)
	srsRepository = repository.NewSrsRepository(db)
}

func setupServices() {
	crsService = service.NewCrsService(epsgRepository, srsRepository)
}

func setupHttpServer() {
	mux := http.NewServeMux()
	server := http_server.NewServer(crsService)

	http_server.RegisterServerHandlers(mux, server)

	log.Println("Starting http server")

	err := http.ListenAndServe(":"+viper.GetString("server.port"), mux)
	if err != nil {
		log.Fatalf("Failed to start http server, %s", err)
	}
}
