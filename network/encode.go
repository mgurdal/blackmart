package network

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Encode turns a go object into a byte array
func Encode(obj interface{}) bytes.Buffer {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(obj)
	if err != nil {
		log.Fatal("Could not encode")
	}

	return buffer

}
