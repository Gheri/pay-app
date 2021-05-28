package service

import "errors"

type User struct {
	Name           string  `bson:"name"`
	EmailId        string  `bson:"emailId"`
	CreditLimit    float64 `bson:"creditLimit"`
	CrrCreditLimit float64 `bson:"crrCreditLimit"`
	Dues           float64 `bson:"dues"`
}

func NewUser(name string, emailId string, creditLimit float64) User {
	return User{name, emailId, creditLimit, creditLimit, 0}
}

func (u *User) Payback(amt float64) error {
	if amt <= 0 {
		return errors.New("Payback failed as amt is zero or negative")
	}
	u.CrrCreditLimit += amt
	u.Dues -= amt
	return nil
}

func (u *User) Purchase(amt float64) error {
	if amt > u.CrrCreditLimit {
		return errors.New("Purchase failed due to credit limit")
	}
	if amt <= 0 {
		return errors.New("Purchase failed as amt is zero or negative")
	}
	u.CrrCreditLimit -= amt
	u.Dues += amt
	return nil
}
