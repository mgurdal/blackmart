package user

import (
	"testing"

	"github.com/mgurdal/blackmarkt/errors"
	"github.com/mgurdal/blackmarkt/market"

	"github.com/mgurdal/blackmarkt/inventory"
)

func TestUserMoveToMarket(t *testing.T) {

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

func TestUserPurchase(t *testing.T) {

	t.Run("PurchaseItem adds the market item to user's inventory", func(t *testing.T) {

		seller := New("Seller")
		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 1,
			Status:   "created",
		}
		seller.Inventory.Add(item)
		seller.MoveToMarket(item, 1, 1)

		buyer := New("Buyer")
		money := &inventory.Item{Name: "Money", Quantity: 10}
		buyer.Inventory.Add(money)

		productname, amount := "TestItem", 1
		buyer.PurchaseItem(seller, productname, amount)

		expected := 1
		if buyer.Inventory.Items[item.Name].Quantity != expected {
			t.Errorf("Expected total inventory count to be %d", expected)
		}
	})
	t.Run("PurchaseItem removes from the seller's market", func(t *testing.T) {

		seller := New("Seller")
		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 1,
			Status:   "created",
		}
		seller.Inventory.Add(item)
		seller.MoveToMarket(item, 1, 1)

		buyer := New("Buyer")
		money := &inventory.Item{Name: "Money", Quantity: 10}
		buyer.Inventory.Add(money)

		buyer.PurchaseItem(seller, item.Name, 1)

		product := seller.Market.Products[0]
		if product.Quantity != 0 {
			t.Error("Expected product quantity to be 0")
		}
	})
	t.Run("PurchaseItem returns error if user does not have enough money", func(t *testing.T) {

		seller := New("Seller")
		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 1,
			Status:   "created",
		}
		seller.Inventory.Add(item)
		seller.MoveToMarket(item, 1, 1)

		buyer := New("Buyer")

		productname, amount := "TestItem", 1
		err := buyer.PurchaseItem(seller, productname, amount)
		if err != errors.NotEnoughMoney {
			t.Errorf("Expected to get %v error", errors.NotEnoughMoney)
		}
	})
	t.Run("PurchaseItem withdraws expected amount of money from buyer", func(t *testing.T) {

		seller := New("Seller")
		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 4,
			Status:   "created",
		}
		seller.Inventory.Add(item)

		productAmount, productPrice := 4, 1
		seller.MoveToMarket(item, productPrice, productAmount)

		buyer := New("Buyer")
		money := &inventory.Item{Name: "Money", Quantity: 10}
		buyer.Inventory.Add(money)

		productName, puchaseAmount := "TestItem", 4

		buyer.PurchaseItem(seller, productName, puchaseAmount)
		expectedBalance := 6
		if buyer.Balance() != expectedBalance {
			t.Errorf(
				"Expected the new balance to be %d got %d",
				expectedBalance,
				buyer.Balance(),
			)
		}
	})
	t.Run("PurchaseItem adds expected amount of money to sellers inventory", func(t *testing.T) {

		seller := New("Seller")
		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 4,
			Status:   "created",
		}
		seller.Inventory.Add(item)

		productAmount, productPrice := 4, 1
		seller.MoveToMarket(item, productPrice, productAmount)

		buyer := New("Buyer")
		money := &inventory.Item{Name: "Money", Quantity: 10}
		buyer.Inventory.Add(money)

		productName, puchaseAmount := "TestItem", 4

		buyer.PurchaseItem(seller, productName, puchaseAmount)
		expectedBalance := 4
		if seller.Balance() != expectedBalance {
			t.Errorf(
				"Expected the new balance to be %d got %d",
				expectedBalance,
				seller.Balance(),
			)
		}
	})
}
