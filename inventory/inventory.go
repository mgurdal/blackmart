package inventory

import "errors"

type Inventory struct {
	Items map[string]*Item
	Total int
}

// Add upserts the target item to the inventory.
func (inv *Inventory) Add(item *Item) {
	items := inv.Items
	if inventoryItem, contains := items[item.Name]; contains {
		inventoryItem.Quantity += item.Quantity
		inv.Total += item.Quantity
	} else {
		items[item.Name] = item
	}
}

// Check compares the consumption requirements with the user inventory
// in order to check the productability of a unit of the target item.
// TODO: we might want to move this method to factory since the production
// task is only related to the factory.
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

// Consume reduces the target item's quantity if its in the inventory.
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
