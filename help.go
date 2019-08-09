package main

import (
	"fmt"
)

// Get the help text for a specific issue
func help(topic string) {
	switch topic {
	case "shine":
		fmt.Println(`Usage: 
  shine [option] [file]	run or edit a shinefile

Options:
  --help [topic] 	display help menu for shine or a specific topic
  --create [file] 	create a new shinefile
  --edit <file> 	edit an existing shinefile`)

	default:
		fmt.Println("Unknown topic:", topic)

	}
}
