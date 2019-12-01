package inventory

import "errors"

type Inventory struct {
	Items map[string]*Item
	q     int
}

func (inv *Inventory) Add(item *Item) {
	items := inv.Items
	if inventoryItem, contains := items[item.Name]; contains {
		inventoryItem.Quantity += item.Quantity
		inv.q += item.Quantity
	} else {
		items[item.Name] = item
	}
}

// Check
func (inv *Inventory) Check(consumptions map[string]*Item) (ok bool) {
	ok = true
	items := inv.Items
	for _, item := range consumptions {
		if inventoryItem, contains := items[item.Name]; contains {
			if inventoryItem.Quantity < item.Quantity {
				ok = false
				break
			}
		} else {
			ok = false
			break
		}
	}

	return
}

// Consume
func (inv *Inventory) Consume(item Item) (err error) {
	items := inv.Items
	inventoryItem, exists := items[item.Name]
	if !exists {
		return errors.New("Item not found")
	}
	if inventoryItem.Quantity < item.Quantity {
		return errors.New("Insufficient quantity")
	}
	inventoryItem.Quantity -= item.Quantity
	return
}
