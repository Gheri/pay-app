package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerchant(t *testing.T) {
	t.Run("Purchase calculates the discount earned", func(t *testing.T) {
		merchant := NewMerchant("m1", "m1@abc.com", 10)
		err := merchant.Purchase(1000)
		assert.NoError(t, err, "Error found while purchase")
		assert.Equal(t, 100.0, merchant.DiscountEarned)
	})
	t.Run("Purchase calculates the discount earned for multiple purchases", func(t *testing.T) {
		merchant := NewMerchant("m1", "m1@abc.com", 10)
		err := merchant.Purchase(1000)
		assert.NoError(t, err, "Error found while purchase")
		assert.Equal(t, 100.0, merchant.DiscountEarned)
		err = merchant.Purchase(1000)
		assert.NoError(t, err, "Error found while purchase")
		assert.Equal(t, 200.0, merchant.DiscountEarned)
	})
	t.Run("Purchase should return error if amt is negative", func(t *testing.T) {
		merchant := NewMerchant("m1", "m1@abc.com", 10)
		err := merchant.Purchase(-100)
		assert.Error(t, err, "No error found while purchase negative amt")
	})
	t.Run("Purchase should return error if amt is zero", func(t *testing.T) {
		merchant := NewMerchant("m1", "m1@abc.com", 10)
		err := merchant.Purchase(0)
		assert.Error(t, err, "No error found while purchase zero amt")
	})
}
