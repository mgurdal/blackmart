package market

import "errors"

var (
	ItemNotFound         = errors.New("Item Not Found")
	InsufficientQuantity = errors.New("Insufficient Quantity")
)

type Market struct {
	Products []*Product
}

// Consume reduces the target item's quantity if its in the inventory.
func (m *Market) Consume(name string, quantity int) error {
	for _, product := range m.Products {
		if product.ItemName == name {
			if product.Quantity < quantity {
				return InsufficientQuantity
			}
			product.Quantity -= quantity
			return nil
		}
	}
	return ItemNotFound

}
