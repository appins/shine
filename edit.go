package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func edit(filename string) error {
	// Read file
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal json
	var shine fileFormat
	err = json.Unmarshal(fileData, &shine)
	if err != nil {
		return errors.New("Shinefile doesn't contain vaild json")
	}

	// Loop until terminated or returning error
	var option int
	for {
		option, _ = multipleChoice("Edit file", []string{
			"Add a new row",
			"Edit an existing row",
			"Remove an existing row",
			"Move a row",
			"View rows",
			"Save and exit",
			"Exit without saving",
		})
		switch option {
		// Add new row
		case 0:
			// Get name, folder, type of row, and data
			name, _ := prompt("Name for row: ")
			if name == "undefined" {
				continue
			}
			folder, _ := prompt("Folder for row (empty = root): ")
			if folder == "" {
				folder = "root"
			}
			if folder == "undefined" {
				continue
			}
			methodInt, _ := multipleChoice("Action", []string{
				"Run a command",
				"Change folders",
				"Exit the shinefile",
			})
			if methodInt == -1 {
				continue
			}
			// Convert the list choice into the right method
			method := []string{"run", "cd", "exit"}[methodInt]

			var data string
			switch method {
			case "cd":
				data, _ = prompt("Folder to open: ")
			case "run":
				data, _ = prompt("Command to run: ")
			}

			shine.Steps = append(shine.Steps, step{name, folder, method, data})

		// Edit an existing row
		case 1:
			var rows []string
			var steps []step
			for _, j := range shine.Steps {
				rows = append(rows, j.Folder+"/"+j.Text)
				steps = append(steps, j)
			}
			chosen, _ := multipleChoice("Existing Rows", rows)
			if chosen == -1 {
				break
			}

			shine.Steps[chosen].edit()
		// View rows
		case 4:
			var rows []string
			for _, j := range shine.Steps {
				rows = append(rows, j.Folder+"/"+j.Text+" - "+j.Method+" "+j.Data)
			}
			multipleChoice("Existing Rows", rows)

		// Save and exit
		case 5:
			saveJson, err := json.MarshalIndent(shine, "\n", "\t")
			if err != nil {
				return err
			}

			return ioutil.WriteFile(filename, saveJson, 0666)
		// Exit without saving
		case 6:
			return nil

		}

	}

	return nil
}
