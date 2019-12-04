package user

import (
	"errors"

	"github.com/mgurdal/blackmarkt/inventory"
	"github.com/mgurdal/blackmarkt/market"
)

type User struct {
	ID int
	*market.Market
	*inventory.Inventory
}

func (u *User) MoveToMarket(item *inventory.Item, price int, amount int) error {
	if item.Quantity >= amount {
		item.Quantity -= amount
		product := &market.Product{
			Quantity: amount,
			ItemName: item.Name,
			Price:    price,
			OwnerID:  u.ID,
		}
		u.Market.Products = append(u.Market.Products, product)
	} else {
		return errors.New("Insufficient Item Amount")
	}
	return nil
}
