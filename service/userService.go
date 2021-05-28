package service

import (
	"github.com/pkg/errors"
)

type UserService struct {
	database Database
}

func (userService *UserService) GetUser(name string) (*User, error) {
	user := userService.database.GetUser(name)
	if user == nil {
		return nil, errors.New("User Not Found !!!")
	}
	return user, nil
}

func (userService *UserService) SaveNewUser(name string, emailId string, creditLimit float64) error {
	user := NewUser(name, emailId, creditLimit)
	err := userService.database.SaveUser(user)
	if err != nil {
		return errors.Wrap(err, "Save user failed in UserService")
	}
	return nil
}

func (userService *UserService) Purchase(userName string, amt float64) error {
	user, err := userService.GetUser(userName)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in UserService")
	}
	err = user.Purchase(amt)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in UserService")
	}
	err = userService.database.UpdateUser(*user)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in UserService")
	}
	return nil
}

func (userService *UserService) Payback(userName string, amt float64) error {
	user, err := userService.GetUser(userName)
	if err != nil {
		return errors.Wrap(err, "Purchase failed in UserService")
	}
	err = user.Payback(amt)
	if err != nil {
		return errors.Wrap(err, "Payback Failed in User Service")
	}
	err = userService.database.UpdateUser(*user)
	if err != nil {
		return errors.Wrap(err, "Payback Failed in User Service")
	}
	return nil
}

func (userService *UserService) GetUsersReachedCreditLimit() ([]string, error) {
	users, err := userService.database.QueryToGetUsersReachedLimits()
	err = errors.Wrap(err, "Query to get user-reached-limit failed")
	return users, err
}

func (userService *UserService) GetAllUsersDues() ([]User, error) {
	users, err := userService.database.QueryToGetUserswithDues()
	err = errors.Wrap(err, "Query to get-all-dues failed")
	return users, err
}
