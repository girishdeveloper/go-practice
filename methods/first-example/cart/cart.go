package cart

import (
	"errors"
	"example/product"
	"os/user"
	"time"

	"github.com/Rhymond/go-money"
)

type Item struct {
	product.Product
	Quantity uint8
}

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	LockedAt  time.Time
	user.User
	Items       []Item
	CountryCode string
	isLocked    bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	return money.New(55, money.INR), nil
}

func (c *Cart) Lock() error {
	if c.isLocked {
		return errors.New("Cart is already locked")
	}
	c.isLocked = true
	c.LockedAt = time.Now()
	return nil
}

func (c *Cart) delete() error {
	return nil
}
