package market

import (
	"testing"

	"github.com/google/uuid"
)

func TestMarket(t *testing.T) {

	t.Run("Consume returns err if the quantity is insufficient", func(t *testing.T) {
		market := &Market{
			Products: []*Product{
				&Product{
					ItemName: "TestItem",
					Quantity: 10,
					Price:    10,
					OwnerID:  uuid.New(),
				},
			},
		}

		name, quantity := "TestItem", 11
		err := market.Consume(name, quantity)
		if err != InsufficientQuantity {
			t.Errorf("Expected to get %v error", InsufficientQuantity)
		}

	})

	t.Run("Consume returns err if the item not in market", func(t *testing.T) {
		market := &Market{
			Products: []*Product{}, // empty market
		}

		name, quantity := "TestItem", 0
		err := market.Consume(name, quantity)
		if err != ItemNotFound {
			t.Errorf("Expected to get %v error", ItemNotFound)
		}

	})

	t.Run("Consume reduces item quantity correctly", func(t *testing.T) {
		product := &Product{
			ItemName: "TestItem",
			Quantity: 10,
			Price:    10,
			OwnerID:  uuid.New(),
		}
		market := &Market{
			Products: []*Product{
				product,
			},
		}

		name, quantity := "TestItem", 9
		market.Consume(name, quantity)

		expectedQuantity := 1

		if product.Quantity != expectedQuantity {
			t.Errorf("Expected item quantity to be %d", expectedQuantity)
		}

	})
}
