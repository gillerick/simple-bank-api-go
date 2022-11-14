package simple_bank_account

import (
	"log"
	"net/http"
	"simple-bank-account/database"
	"simple-bank-account/postgres"
	"simple-bank-account/service"
)

func main() {
	service := Service{}

	if err := service.
}

type Service struct {
	HttpServer *http.Server
}

func (s *Service) Run() {
	var err error

	//Setup a postgres database connection
	//TODo: Provide database configs
	pgDb, err := postgres.NewConnection()
	if err != nil {
		log.Fatal("could not establish connection with the database")
	}

	//Setup database
	database := database.NewDatabaseHandler(pgDb)

	//Setup services
	//ToDo: Inject repository into services
	accountService := service.NewAccountService()
	customerService := service.NewCustomerService()

}
