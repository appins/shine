package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get command line arguments, if none, load "shinefile"
	args := os.Args[1:]
	if len(args) == 0 {
		dat, err := ioutil.ReadFile("shinefile")
		if err != nil {
			fmt.Fprintln(os.Stderr, "No shinefile found")
			os.Exit(1)
		}

		// Display menu
		err = disp(dat)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}

}
