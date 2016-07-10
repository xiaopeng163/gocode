package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(s.Servers[0].ServerIP)

	s1 := Server{"Beijing1", "1.1.1.1"}
	s2 := Server{"Shanghai", "2.2.2.2"}
	a := []Server{s1, s2}
	sl := Serverslice{a}
	fmt.Println(sl)

}
