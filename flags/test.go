package main

import (
	"fmt"
	"os"
	"runtime"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		ConfigFile    string `short:"f" long:"config-file" description:"specifying a config file"`
		DisableStdlog bool   `long:"disable-stdlog" description:"disable standard logging"`
		CPUs          int    `long:"cpus" default:"1" description:"specify the number of CPUs to be used"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println(opts.CPUs)
}
