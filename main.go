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

		// Display shinefile
		err = disp(dat)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	} else if len(args) == 1 {
		// Check if the argument is a known option
		switch args[0] {
		case "--create", "-c":

		case "--help", "-h":

		default:
			// Check if the argument is an unknown option
			if args[0][0] == '-' {
				fmt.Fprintln(os.Stderr, args[0], "is an unrecognized option")
				os.Exit(1)
			}

			// Otherwise, assume it's a filename
			dat, err := ioutil.ReadFile(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, args[0], "is not a valid shinefile")
				os.Exit(1)
			}

			// Display shinefile
			err = disp(dat)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		}
	}

}
