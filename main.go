package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//	ui "github.com/gizak/termui/v3"
)

type step struct {
	name   string
	method string
	data   string
}

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

	// if err := ui.init(); err != nil {
	// 	log.Fatal("termui err: ", err)
	//	fmt.Fprintln(os.Stderr, "error: ", err)
	// }
	// defer ui.Close()

}
