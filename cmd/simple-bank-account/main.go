package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"simple-bank-account/configs"
	"simple-bank-account/controllers"
	"simple-bank-account/database"
	"simple-bank-account/http"
	"simple-bank-account/repositories"
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

	//Setup a database connection
	pgDb, err := database.NewConnection(config.DB)
	if err != nil {
		log.Fatal("could not establish connection with the database")
	}

	//Set up database
	dbHandler := repositories.NewDatabaseHandler(pgDb)

	//Run database migrations
	err = dbHandler.Migrate()
	if err != nil {
		return fmt.Errorf("repositories migrations failed: %w", err)
	}

	//Set up repositories
	customerRepository := repositories.NewCustomerRepository(dbHandler)
	accountRepository := repositories.NewAccountRepository(dbHandler)
	cardRepository := repositories.NewCardRepository(dbHandler)

	//Setup services
	customerService := services.NewCustomerService(customerRepository)
	accountService := services.NewAccountsService(accountRepository)
	cardService := services.NewCardService(cardRepository)

	//Set up HTTP handler and router
	customerHandler := controllers.NewCustomerController(*customerService)
	accountHandler := controllers.NewAccountsController(*accountService)
	cardHandler := controllers.NewCardController(*cardService)

	//Initialize server
	s.HttpServer = http.NewServer(config.App, customerHandler, accountHandler, cardHandler)

	//Start the HTTP handler
	s.HttpServer.Run()

	// Wait for OS termination signal
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	return nil

}
