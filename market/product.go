package market

import "github.com/google/uuid"

type Product struct {
	Price    int
	Quantity int
	OwnerID  uuid.UUID
	ItemName string // Item.Name
}
