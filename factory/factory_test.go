package factory

import (
	"testing"
	"time"

	"github.com/mgurdal/blackmarkt/inventory"
	"github.com/mgurdal/blackmarkt/user"
)

func TestFactoryUpdate(t *testing.T) {

	t.Run("Update resets the updated at if deposit is full", func(t *testing.T) {
		now := time.Now()
		tenSecondsBefore := now.Add(time.Second * time.Duration(-10))
		factory := &Factory{
			ItemName:  "TestItem",
			Speed:     1,
			Deposit:   1,
			Limit:     1,
			UpdatedAt: tenSecondsBefore,
		}

		factory.Update(now)

		if factory.UpdatedAt != now {
			t.Errorf(
				"Expected the updated at to be %s got %s",
				now,
				factory.UpdatedAt,
			)
		}
	})

	t.Run("Update - updated at does not change if the update gets triggered before completing a period.", func(t *testing.T) {
		now := time.Now()
		halfASecond := now.Add(time.Millisecond * time.Duration(-500))
		factory := &Factory{
			ItemName:  "TestItem",
			Speed:     1,
			Deposit:   1,
			Limit:     10,
			UpdatedAt: halfASecond,
		}

		factory.Update(now)

		if factory.UpdatedAt != halfASecond {
			t.Errorf(
				"Expected the updated at to be %s got %s",
				now,
				factory.UpdatedAt,
			)
		}
	})

	// "Update resets updated at if has not enough raw materials"
	t.Run("Update produces items correctly if periods passes", func(t *testing.T) {
		now := time.Now()
		fiveSecondsAgo := now.Add(time.Second * time.Duration(-5))

		rawMaterials := map[string]*inventory.Item{
			"Source1": &inventory.Item{
				Name:     "Source1",
				Quantity: 2,
			},
			"Source2": &inventory.Item{
				Name:     "Source2",
				Quantity: 1,
			},
		}
		inventoryItems := map[string]*inventory.Item{
			"Source1": &inventory.Item{
				Name:     "Source1",
				Quantity: 15,
			},
			"Source2": &inventory.Item{
				Name:     "Source2",
				Quantity: 10,
			},
		}
		user := &user.User{
			Inventory: &inventory.Inventory{
				Items: inventoryItems,
			},
		}
		factory := &Factory{
			ItemName:     "TestItem",
			Speed:        1,
			Deposit:      0,
			Limit:        10,
			UpdatedAt:    fiveSecondsAgo,
			User:         user,
			RawMaterials: rawMaterials,
		}

		factory.Update(now)
		expectedDeposit := 5
		if factory.Deposit != expectedDeposit {
			t.Errorf(
				"Expected the deposit to be %d got %d ",
				expectedDeposit,
				factory.Deposit,
			)
		}

		expectedSource1Quantity := 5
		source1Quantity := user.Inventory.Items["Source1"].Quantity
		if source1Quantity != expectedSource1Quantity {
			t.Errorf(
				"Expected the Source1 to be %d got %d ",
				expectedSource1Quantity,
				source1Quantity,
			)
		}

		expectedSource2Quantity := 5
		source2Quantity := user.Inventory.Items["Source2"].Quantity
		if source2Quantity != expectedSource2Quantity {
			t.Errorf(
				"Expected the Source2 to be %d got %d ",
				expectedSource2Quantity,
				source2Quantity,
			)
		}

		if factory.UpdatedAt != now {
			t.Errorf(
				"Expected the updated at to be %s got %s",
				now,
				factory.UpdatedAt,
			)
		}
	})

	t.Run("Update resets the updated at if periods passes", func(t *testing.T) {
		now := time.Now()
		fiveAndHalfSecondsAgo := now.Add(time.Millisecond * time.Duration(-5500))
		rawMaterials := map[string]*inventory.Item{
			"Source1": &inventory.Item{
				Name:     "Source1",
				Quantity: 2,
			},
			"Source2": &inventory.Item{
				Name:     "Source2",
				Quantity: 1,
			},
		}
		inventoryItems := map[string]*inventory.Item{
			"Source1": &inventory.Item{
				Name:     "Source1",
				Quantity: 15,
			},
			"Source2": &inventory.Item{
				Name:     "Source2",
				Quantity: 10,
			},
		}
		user := &user.User{
			Inventory: &inventory.Inventory{
				Items: inventoryItems,
			},
		}
		factory := &Factory{
			ItemName:     "TestItem",
			Speed:        1,
			Deposit:      1,
			Limit:        10,
			UpdatedAt:    fiveAndHalfSecondsAgo,
			RawMaterials: rawMaterials,
			User:         user,
		}

		factory.Update(now)
		want := now.Add(time.Millisecond * time.Duration(-500))
		if factory.UpdatedAt != want {
			t.Errorf(
				"Expected the updated at to be %s got %s",
				want,
				factory.UpdatedAt,
			)
		}
	})
}

func TestFactoryCollect(t *testing.T) {

	t.Run("Collect transports the expected amount of item to user inventory", func(t *testing.T) {

		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 0,
		}
		inventoryItems := map[string]*inventory.Item{
			"TestItem": item,
		}

		user := &user.User{
			Inventory: &inventory.Inventory{
				Items: inventoryItems,
			},
		}
		factory := &Factory{
			ItemName: "TestItem",
			Speed:    1,
			Deposit:  3,
			Limit:    10,
			User:     user,
		}

		factory.Collect()

		expectedQuantity := 3
		userQuantity := user.Inventory.Items["TestItem"].Quantity
		if userQuantity != expectedQuantity {
			t.Errorf(
				"Expected the item quantity to be %d got %d",
				expectedQuantity,
				userQuantity,
			)
		}
	})

	t.Run("Collect resets the factory deposit", func(t *testing.T) {

		item := &inventory.Item{
			Name:     "TestItem",
			Quantity: 0,
		}
		inventoryItems := map[string]*inventory.Item{
			"TestItem": item,
		}

		user := &user.User{
			Inventory: &inventory.Inventory{
				Items: inventoryItems,
			},
		}
		factory := &Factory{
			ItemName: "TestItem",
			Speed:    1,
			Deposit:  3,
			Limit:    10,
			User:     user,
		}

		factory.Collect()

		if factory.Deposit != 0 {
			t.Error(
				"Expected the factory deposit to be 0.",
			)
		}
	})
}
