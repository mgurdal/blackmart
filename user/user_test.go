package user

import (
	"testing"

	"github.com/google/uuid"
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
	buyerid := uuid.New()
	buyer := User{
		ID: buyerid,
		Inventory: &inventory.Inventory{
			Items: map[string]*inventory.Item{
				"TestItem": &inventory.Item{Name: "TestItem", Quantity: 1},
			},
		},
		Market: &market.Market{
			Products: []*market.Product{},
		},
	}
	sellerid := uuid.New()
	product := &market.Product{
		Price:    1,
		Quantity: 1,
		OwnerID:  sellerid,
		ItemName: "TestItem",
	}
	seller := User{
		ID: sellerid,
		Inventory: &inventory.Inventory{
			Items: map[string]*inventory.Item{},
		},
		Market: &market.Market{
			Products: []*market.Product{
				product,
			},
		},
	}

	t.Run("PurchaseItem adds the market item to user's inventory", func(t *testing.T) {

		productname, amount := "TestItem", 1
		buyer.PurchaseItem(seller.Market, productname, amount)

		if buyer.Inventory.Total != 1 {
			t.Errorf("Expected total inventory count to be %d", buyer.Inventory.Total)
		}
	})

	t.Run("PurchaseItem removes from the seller's market", func(t *testing.T) {

		productname, amount := "TestItem", 1
		buyer.PurchaseItem(seller.Market, productname, amount)

		if product.Quantity != 0 {
			t.Error("Expected product quantity to be 0")
		}
	})

}
