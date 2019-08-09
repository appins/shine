package main

import (
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

	return nil

}
