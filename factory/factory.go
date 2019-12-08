package factory

import (
	"log"

	"time"

	"github.com/google/uuid"
	"github.com/mgurdal/blackmarkt/inventory"
	"github.com/mgurdal/blackmarkt/user"
)

type Factory struct {
	ID           uuid.UUID
	ItemName     string
	Speed        int
	Deposit      int
	Limit        int
	UpdatedAt    time.Time
	Status       string
	RawMaterials map[string]*inventory.Item // ?
	*user.User
}

func (f *Factory) Start() {
	f.Status = "running"
	f.UpdatedAt = time.Now()
	go func(f *Factory) {
		for {
			<-time.After(1 * time.Second)
			f.Update(time.Now())
		}
	}(f)
}

// IsFull check whether the deposit is reached to the limit.
func (f *Factory) IsFull() bool {
	if f.Deposit < f.Limit {
		return false
	}
	return true
}

func (f *Factory) Update(now time.Time) {
	delta := time.Since(f.UpdatedAt) / time.Second
	if f.IsFull() {
		f.UpdatedAt = now
		log.Println("Deposit is full.")
		return
	}
	periods := int(delta) * f.Speed
	for i := 0; i < periods; i++ {

		hasEnoughItems := f.User.Inventory.Check(f.RawMaterials)
		if !hasEnoughItems {
			f.UpdatedAt = time.Now()
			log.Println("Low on raw materials.")
			break
		}

		for _, item := range f.RawMaterials {
			f.User.Inventory.Consume(item.Name, item.Quantity)
		}

		f.Deposit += 1
		log.Printf("Deposit has increased to %d", f.Deposit)
		f.UpdatedAt = f.UpdatedAt.Add(time.Duration(f.Speed) * time.Second)
	}

}

// Collect transports factory deposit to user inventory.
func (f *Factory) Collect() {
	f.User.Inventory.Items[f.ItemName].Quantity += f.Deposit
	// reset factory deposit
	f.Deposit = 0
}
