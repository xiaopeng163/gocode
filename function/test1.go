package main

import "strings"
import "fmt"

// BGPAttrFlag is uint8
type BGPAttrFlag uint8

const (
	BGP_ATTR_FLAG_EXTENDED_LENGTH BGPAttrFlag = 1 << 4
	BGP_ATTR_FLAG_PARTIAL         BGPAttrFlag = 1 << 5
	BGP_ATTR_FLAG_TRANSITIVE      BGPAttrFlag = 1 << 6
	BGP_ATTR_FLAG_OPTIONAL        BGPAttrFlag = 1 << 7
)

func (f BGPAttrFlag) String() string {
	strs := make([]string, 0, 4)
	if f&BGP_ATTR_FLAG_EXTENDED_LENGTH > 0 {
		strs = append(strs, "EXTENDED_LENGTH")
	}
	if f&BGP_ATTR_FLAG_PARTIAL > 0 {
		strs = append(strs, "PARTIAL")
	}
	if f&BGP_ATTR_FLAG_TRANSITIVE > 0 {
		strs = append(strs, "TRANSITIVE")
	}
	if f&BGP_ATTR_FLAG_OPTIONAL > 0 {
		strs = append(strs, "OPTIONAL")
	}
	return strings.Join(strs, "|")
}

func main() {
	var flag BGPAttrFlag = 0xc0
	fmt.Println(flag)
}
