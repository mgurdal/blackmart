package user

import (
	"testing"

	"github.com/mgurdal/blackmarkt/market"

	"github.com/mgurdal/blackmarkt/inventory"
)

func TestUser(t *testing.T) {

	t.Run("MoveToMarket returns error if the requested item quantity is not available in the user inventory", func(t *testing.T) {

		item := &inventory.Item{Name: "TestItem", Quantity: 1}
		inv := &inventory.Inventory{
			Items: map[string]*inventory.Item{
				item.Name: item,
			},
		}
		market := &market.Market{
			Products: []*market.Product{},
		}
		user := &User{Inventory: inv, Market: market}

		price, amount := 11, 5
		err := user.MoveToMarket(item, price, amount)

		if err == nil {
			t.Errorf("Expected to get Insufficient Item Amount error")
			return
		}
	})

	t.Run("MoveToMarket adds the item to the market with given price and amount.", func(t *testing.T) {

		item := &inventory.Item{Name: "TestItem", Quantity: 40}
		inv := &inventory.Inventory{
			Items: map[string]*inventory.Item{
				item.Name: item,
			},
		}
		market := &market.Market{
			Products: []*market.Product{},
		}
		user := &User{Inventory: inv, Market: market}

		price, amount := 11, 5
		user.MoveToMarket(item, price, amount)

		if len(user.Market.Products) != 1 {
			t.Errorf("Expected the product count to be 1.")
			return
		}
		product := user.Market.Products[0]
		if product.Quantity != amount {
			t.Errorf("Expected the product quantity to be %d", amount)
			return
		}
		if product.Price != price {
			t.Errorf("Expected the product price to be %d", price)
			return
		}
	})
}
