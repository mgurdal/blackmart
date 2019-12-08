package errors

import "errors"

var (
	ItemNotFound         = errors.New("Item Not Found")
	InsufficientQuantity = errors.New("Insufficient Quantity")
	NotEnoughMoney       = errors.New("Not Enough Money")
)
