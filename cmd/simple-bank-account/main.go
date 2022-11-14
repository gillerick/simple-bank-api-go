package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"simple-bank-account/configs"
	"simple-bank-account/database"
	"simple-bank-account/http"
	"simple-bank-account/postgres"
	"simple-bank-account/services"
	"syscall"
)

func main() {
	service := Service{}

	if err := service.Run(); err != nil {
		log.Fatalf("unable to start services %s", err)
	}
}

type Service struct {
	HttpServer *http.Server
}

func (s *Service) Run() error {
	var err error

	// Fetch app configurations. Empty paths reads configs from a set default path
	ymlConfig := configs.ReadYaml("")
	config := configs.GetConfig(*ymlConfig)

	//Setup a postgres database connection
	//TODo: Provide database configs
	pgDb, err := postgres.NewConnection(config.DB)
	if err != nil {
		log.Fatal("could not establish connection with the database")
	}

	//Setup database
	database := database.NewDatabaseHandler(pgDb)

	//Run database migrations
	err = database.Migrate(pgDb)
	if err != nil {
		return fmt.Errorf("database migrations failed: %w", err)
	}

	//Setup services
	//ToDo: Inject repository into services
	accountService := services.NewAccountService()
	customerService := services.NewCustomerService()

	//Set up HTTP handler and router
	s.HttpServer = http.NewServer(config.App)

	//Start the HTTP handler
	s.HttpServer.Run()

	// Wait for OS termination signal
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	return nil

}
