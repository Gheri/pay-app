package service

import "errors"

type Merchant struct {
	Name               string  `bson:"name"`
	EmailId            string  `bson:"emailId"`
	DiscountPercentage float64 `bson:"discountPercentage"`
	DiscountEarned     float64 `bson:"discountEarned"`
}

func NewMerchant(name string, emailId string, discountPercentage float64) Merchant {
	return Merchant{name, emailId, discountPercentage, 0}
}

func (merchant *Merchant) Purchase(amt float64) error {
	if amt <= 0 {
		return errors.New("amt cannot be zero or negative")
	}
	discountEarned := (amt * merchant.DiscountPercentage) / 100.0
	merchant.DiscountEarned += discountEarned
	return nil
}
