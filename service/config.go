package service

func GetMongoString() string {
	return "mongodb://localhost:27017"
}

func GetUsersCollectionName() string {
	return "users"
}

func GetMerchantsCollectionName() string {
	return "merchants"
}

func GetTransactionsCollectionName() string {
	return "transactions"
}

func GetDatabaseName() string {
	return "test"
}
