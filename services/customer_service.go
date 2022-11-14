package services

type CustomerService struct {
	repository DataStore
}

func NewCustomerService(repository DataStore) *CustomerService {
	return &CustomerService{repository: repository}
}
