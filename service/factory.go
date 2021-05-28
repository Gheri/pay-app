package service

var database Database = NewDatabase()
var userService *UserService = &UserService{database: database}
var merchantService *MerchantService = &MerchantService{database: database}
var transactionService *TransactionService = &TransactionService{userService, merchantService}

func GetUserService() *UserService {
	return userService
}
func GetMerchantService() *MerchantService {
	return merchantService
}
func GetTransactionService() *TransactionService {
	return transactionService
}
