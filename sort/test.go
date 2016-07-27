package main

import "fmt"
import "sort"

func main() {
	s := []string{"Jim", "Mark", "John"}

	fmt.Println(s)
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
