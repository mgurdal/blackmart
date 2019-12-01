package inventory

import "testing"

func TestInventory(t *testing.T) {

	t.Run("Add sums the quantities if item exists", func(t *testing.T) {
		item := &Item{
			"TestItem",
			10,
			"test",
		}
		items := map[string]*Item{
			item.Name: item,
		}
		inv := Inventory{
			Items: items,
		}
		inv.Add(item)

		expectedQuantity := 20
		updatedItem := inv.Items[item.Name]
		if updatedItem.Quantity != expectedQuantity {
			t.Errorf(
				"Failed to add item: got %d want %d",
				updatedItem.Quantity,
				expectedQuantity,
			)
		}
	})

	t.Run("Add the item in inventory if does not exists", func(t *testing.T) {
		item := &Item{
			"TestItem",
			10,
			"test",
		}
		inv := Inventory{
			Items: map[string]*Item{},
		}
		inv.Add(item)

		_, itemAdded := inv.Items[item.Name]
		if !itemAdded {
			t.Errorf(
				"Failed to add item %s",
				item.Name,
			)
		}
	})

	t.Run("Check returns true if all items viable", func(t *testing.T) {

		items := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			items[name] = &Item{
				name,
				10,
				"test",
			}
		}
		inv := Inventory{Items: items}
		got := inv.Check(items)
		want := true
		if got != want {
			t.Errorf("Expected to get %t got %t", want, got)
		}

	})

	t.Run("Check returns false if the inventory has not enough quantities for an item ", func(t *testing.T) {

		items := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			items[name] = &Item{
				name,
				10,
				"test",
			}
		}
		items["Item1"].Quantity = 0

		consumptions := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			consumptions[name] = &Item{
				name,
				10,
				"test",
			}
		}

		inv := Inventory{Items: items}
		got := inv.Check(consumptions)
		want := false
		if got != want {
			t.Errorf("Expected to get %t got %t", want, got)
		}

	})

	t.Run("Check returns false if the inventory is empty", func(t *testing.T) {

		items := map[string]*Item{}

		consumptions := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			consumptions[name] = &Item{
				name,
				10,
				"test",
			}
		}

		inv := Inventory{Items: items}
		got := inv.Check(consumptions)
		want := false
		if got != want {
			t.Errorf("Expected to get %t got %t", want, got)
		}

	})

	t.Run("Consume reduces the item quantity", func(t *testing.T) {

		items := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			items[name] = &Item{
				name,
				10,
				"test",
			}
		}

		inv := Inventory{Items: items}

		item := Item{
			"Item2",
			7,
			"test",
		}
		inv.Consume(item)

		got := inv.Items[item.Name].Quantity
		want := 3
		if got != want {
			t.Errorf("Expected the item quantity to be %d found %d", want, got)
		}

	})

	t.Run("Consume returns error if item does not exists", func(t *testing.T) {

		items := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			items[name] = &Item{
				name,
				10,
				"test",
			}
		}

		inv := Inventory{Items: items}

		item := Item{Name: "NotExistingItem"}
		err := inv.Consume(item)

		if err == nil {
			t.Errorf("Expected to get an error")
		}

	})
	t.Run("Consume not returns error if item exists", func(t *testing.T) {

		items := map[string]*Item{}
		for _, name := range []string{"Item1", "Item2", "Item3"} {
			items[name] = &Item{
				name,
				10,
				"test",
			}
		}

		inv := Inventory{Items: items}

		item := Item{
			"Item2",
			7,
			"test",
		}
		err := inv.Consume(item)

		if err != nil {
			t.Errorf("Expected to not get an error")
		}

	})

	t.Run("Consume not returns error has not enough quantity", func(t *testing.T) {

		items := map[string]*Item{
			"TestItem": &Item{
				"TestItem",
				10,
				"test",
			},
		}

		inv := Inventory{Items: items}

		item := Item{
			"TestItem",
			100,
			"test",
		}
		err := inv.Consume(item)

		if err == nil {
			t.Errorf("Expected to get an error")
		}

	})
}
