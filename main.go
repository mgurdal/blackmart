package main

import (
	"time"

	"github.com/mgurdal/blackmarkt/factory"
	"github.com/mgurdal/blackmarkt/user"

	"github.com/mgurdal/blackmarkt/inventory"
	"github.com/mgurdal/blackmarkt/market"
)

func main() {
	usr := &user.User{
		ID: 123,
		Market: &market.Market{
			Products: []*market.Product{},
		},
		Inventory: &inventory.Inventory{
			Items: map[string]*inventory.Item{},
		},
	}
	newApples := &inventory.Item{
		Name:     "Apple",
		Quantity: 50,
	}
	usr.Inventory.Add(newApples)
	forever := make(chan bool)
	factory := factory.Factory{
		ItemName:  newApples.Name,
		Speed:     1,
		Deposit:   0,
		Limit:     10,
		UpdatedAt: time.Now(),
		Status:    "created",
		RawMaterials: map[string]*inventory.Item{
			newApples.Name: newApples,
		},
		User: usr,
	}
	factory.Start()
	<-forever
}
