package store

import (
	"sync"

	"github.com/google/uuid"
)

var once sync.Once

type Store map[uuid.UUID]interface{}

func (s *Store) Register(ID uuid.UUID, element interface{}) {
	(*s)[ID] = element
}

func (s *Store) Get(ID uuid.UUID) interface{} {
	return (*s)[ID]
}

var (
	store Store
)

// Get the singleton Store instance
func GetStore() *Store {

	once.Do(func() { // <-- atomic, does not allow repeating

		store = make(Store) // <-- thread safe

	})

	return &store
}
