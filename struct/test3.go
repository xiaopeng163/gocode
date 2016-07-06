package main

import "encoding/binary"
import "fmt"

// DefaultRouteDistinguisher
type DefaultRouteDistinguisher struct {
	Type  uint16
	Value []byte
}

// DecodeFromBytes
func (rd *DefaultRouteDistinguisher) DecodeFromBytes(data []byte) error {
	rd.Type = binary.BigEndian.Uint16(data[0:2])
	rd.Value = data[2:8]
	return nil
}

// main function
func main() {

	rd := DefaultRouteDistinguisher{}
	fmt.Println(rd)
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	rd.DecodeFromBytes(data)
	fmt.Println(rd)
}
