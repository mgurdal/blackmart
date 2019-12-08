package market

import "errors"

type Market struct {
	Products []*Product
}

// Consume reduces the target item's quantity if its in the inventory.
func (m *Market) Consume(name string, quantity int) error {
	for _, product := range m.Products {
		if product.ItemName == name {
			if product.Quantity < quantity {
				return errors.New("Insufficient quantity")
			}
			product.Quantity -= quantity
			return nil
		}
	}
	return errors.New("Item not found")

}
