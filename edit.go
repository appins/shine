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
			var st step
			ok := st.Create()
			if ok {
				shine.Steps = append(shine.Steps, st)
			}

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

			shine.Steps[chosen].Edit()
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
