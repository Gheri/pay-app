package service

import (
	"github.com/pkg/errors"
)

type TransactionService struct {
	userService     *UserService
	merchantService *MerchantService
}

func (transactionService *TransactionService) Execute(userName string, merchantName string, amt float64) error {
	err := transactionService.userService.Purchase(userName, amt)
	if err != nil {
		return errors.Wrap(err, "Transaction execute failed in TransactionService")
	}
	err = transactionService.merchantService.Purchase(merchantName, amt)
	if err != nil {
		return errors.Wrap(err, "Transaction execute failed in TransactionService")
	}
	return nil
}
