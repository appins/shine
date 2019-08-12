package main

import (
	"encoding/json"
	"errors"
	"os"
)

func create(filename string) error {
	var fil *os.File

	_, err := os.Stat(filename)
	if err == nil {
		return errors.New("File exists, use --edit instead")
	}

	fil, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer fil.Close()

	title, err := prompt("Title for this shinefile: ")
	if err != nil {
		return err
	}

	var shine fileFormat
	shine.Metadata = map[string]string{"title": title}
	shine.Steps = []step{step{"Exit", "root", "exit", ""}}

	jdat, err := json.MarshalIndent(shine, "\n", "\t")
	if err != nil {
		return err
	}

	fil.Write(jdat)

	return edit(filename)

}
