package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("Purchase/Payback should adjust dues", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(200)
		assert.NoError(t, err, "Error found while purchasing")
		assert.Equal(t, 200.0, user.Dues)
		err = user.Payback(100)
		assert.NoError(t, err, "Error found while payback")
		assert.Equal(t, 100.0, user.Dues)
	})
	t.Run("Purchase/Payback should adjust current credit limit", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(200)
		assert.NoError(t, err, "Error found while purchasing")
		assert.Equal(t, 300.0, user.CrrCreditLimit)
		err = user.Payback(100)
		assert.NoError(t, err, "Error found while payback")
		assert.Equal(t, 400.0, user.CrrCreditLimit)
	})
	t.Run("Payback should return error if amt is negative", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Payback(-100)
		assert.Error(t, err, "No error found while payback negative amt")
	})
	t.Run("Payback should return error if amt is zero", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Payback(0)
		assert.Error(t, err, "No error found while payback zero amt")
	})
	t.Run("Purchase should return error if amt is negative", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(-100)
		assert.Error(t, err, "No error found while purchase negative amt")
	})
	t.Run("Purchase should return error if amt is zero", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(0)
		assert.Error(t, err, "No error found while purchase zero amt")
	})
	t.Run("Purchase should not return error if amt is reached credit limit", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(500)
		assert.NoError(t, err, "Error found while purchase amt equal to credit limit")
	})
	t.Run("Purchase should return error if amt is greater than credit limit", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(600)
		assert.Error(t, err, "No error found while purchase amt greater than credit limit")
	})
	t.Run("Purchase should not return error for multiple purchases if amt is reached credit limit", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(200)
		assert.NoError(t, err, "Error found while purchase amt less than credit limit")
		err = user.Purchase(300)
		assert.NoError(t, err, "Error found while purchase amt equal to credit limit")
	})
	t.Run("Purchase should return error if amt is greater than credit limit", func(t *testing.T) {
		user := NewUser("test1", "t1@abc.com", 500)
		err := user.Purchase(300)
		assert.NoError(t, err, "Error found while purchase amt less than credit limit")
		err = user.Purchase(300)
		assert.Error(t, err, "No error found while purchase amt greater than credit limit")
	})
}
