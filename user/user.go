package user

import (
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	marketErrors "github.com/mgurdal/blackmarkt/errors"
	"github.com/mgurdal/blackmarkt/inventory"
	"github.com/mgurdal/blackmarkt/market"
)

type User struct {
	ID uuid.UUID
	*market.Market
	*inventory.Inventory
	Conn net.Conn
}

// New creates a user with empty market
// and  inventory with 0 money
func New(name string) *User {
	userid := uuid.New()
	money := &inventory.Item{Name: "Money", Quantity: 0}

	usr := &User{
		ID: userid,
		Inventory: &inventory.Inventory{
			Items: map[string]*inventory.Item{
				"Money": money,
			},
		},
		Market: &market.Market{
			Products: []*market.Product{},
		},
	}
	return usr
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

func (u *User) Balance() int {
	money := u.Inventory.Items["Money"].Quantity
	return money
}

func (u *User) Withdraw(amount int) error {
	err := u.Inventory.Consume("Money", amount)
	return err
}

func (u *User) PurchaseItem(seller *User, name string, amount int) error {

	for _, product := range seller.Market.Products {
		if product.ItemName == name {
			if product.Quantity < amount {
				return marketErrors.InsufficientQuantity
			}
			// calculate cost/revenue
			cost := product.Price * amount
			balance := u.Balance()
			if cost > balance {
				return marketErrors.NotEnoughMoney
			}
			err := u.Withdraw(cost)
			if err != nil {
				log.Fatal(err)
				return err
			}

			product.Quantity -= amount

			revenue := &inventory.Item{
				Name:     "Money",
				Quantity: cost,
			}
			seller.Inventory.Add(revenue)

			newItem := &inventory.Item{
				Name:     name,
				Quantity: amount,
			}
			u.Inventory.Add(newItem)
			return nil
		}
	}
	return marketErrors.ItemNotFound
}
