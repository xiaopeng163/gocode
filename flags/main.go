package main

import (
	"flag"
	"fmt"
)

func main() {

	// use pointer
	var infile *string = flag.String("i", "infile",
		"File contains values for sorting")

	// use var
	var outfile string
	flag.StringVar(&outfile, "o", "outfile", "File to receive sorted values")

	var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

	// use address
	var count int
	flag.IntVar(&count, "c", 0, "number of files")

	var debug bool
	flag.BoolVar(&debug, "debug", false, "if use debug")

	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", outfile, "algorithm =", *algorithm, "count =", count, "debug =", debug)
	}
}
