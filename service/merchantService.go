package service

import "github.com/pkg/errors"

type MerchantService struct {
	database Database
}

func (merchantService *MerchantService) SaveNewMerchant(name string, emailId string, discountPercentage float64) error {
	merchant := NewMerchant(name, emailId, discountPercentage)
	err := merchantService.database.SaveMerchant(merchant)
	if err != nil {
		return errors.Wrap(err, "Save new merchant failed in MerchantService")
	}
	return nil
}

func (merchantService *MerchantService) Purchase(name string, amt float64) error {
	merchant := merchantService.database.GetMerchant(name)
	if merchant == nil {
		return errors.New("Merchant not found")
	}
	err := merchant.Purchase(amt)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in MerchantService")
	}
	err = merchantService.database.UpdateMerchant(*merchant)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in MerchantService")
	}
	return nil
}

func (merchantService *MerchantService) GetDiscountPercent(name string) (float64, error) {
	merchant := merchantService.database.GetMerchant(name)
	if merchant == nil {
		return 0, errors.New("Merchant not found")
	}
	return merchant.DiscountPercentage, nil
}
