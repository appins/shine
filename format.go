package main

type step struct {
	Text   string `json:"text"`
	Folder string `json:"folder"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

// Edit dialogue
func (st *step) edit() {
	option, _ := multipleChoice("What field would you like to change", []string{
		"Name of step",
		"Containing folder",
		"Action of step",
	})

	switch option {
	case -1:
		return
	// Edit name
	case 0:
		name, _ := prompt("New name for step: ")
		if name != "undefined" {
			return
		}

		st.Text = name
	// Edit folder
	case 1:
		folder, _ := prompt("New containing folder: ")
		if folder == "undefined" {
			folder = "root"
		}

		st.Folder = folder
	// Edit action
	case 2:
		methodInt, _ := multipleChoice("Method", []string{
			"Run a command",
			"Change Folder",
			"Exit the shinefile",
		})

		if methodInt == -1 {
			return
		}

		method := []string{"run", "cd", "exit"}[methodInt]

		var data string
		switch method {
		case "run":
			data, _ = prompt("Command to run: ")
		case "cd":
			data, _ = prompt("Folder to open: ")
		}

		st.Method = method
		st.Data = data
	}
}

type fileFormat struct {
	Steps    []step            `json:"steps"`
	Metadata map[string]string `json:"meta"`
}
