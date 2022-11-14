package simple_bank_account

import (
	"log"
	"net/http"
	"simple-bank-account/configs"
	"simple-bank-account/database"
	"simple-bank-account/postgres"
	"simple-bank-account/services"
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

	//Setup services
	//ToDo: Inject repository into services
	accountService := services.NewAccountService()
	customerService := services.NewCustomerService()

}
