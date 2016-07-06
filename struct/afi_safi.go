package main

import "fmt"

type AFISAFI int

func (afi_safi AFISAFI) String() string {
	switch afi_safi {
	case 1:
		return "ipv4"
	}
	return fmt.Sprintf("unknown(%d)", afi_safi)
}

func main() {
	var afi_safi AFISAFI
	afi_safi = 1
	fmt.Println(afi_safi.String())
	afi_safi = 2
	fmt.Println(afi_safi.String())

}
