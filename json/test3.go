package main

import "fmt"
import "encoding/json"

func main() {

	type Person struct {
		FName string `json:"First Name"`
		LName string
		Age   int
		Title []string
	}

	p1 := Person{"Peter", "Json", 31, []string{"CEO", "CFO"}}
	fmt.Println(p1)
	// to json format
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(string(bs))
	}

	// load json to person type
	var p2 Person
	err = json.Unmarshal(bs, &p2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p2)
	}

}
