package main

import (
	"encoding/json"
	"errors"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type step struct {
	Text   string `json:"text"`
	Folder string `json:"folder"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

type fileFormat struct {
	Steps    []step            `json:"steps"`
	Metadata map[string]string `json:"meta"`
}

// Display a shinefile
func disp(fileData []byte) error {
	// Unmarshal json
	var shine fileFormat
	err := json.Unmarshal(fileData, &shine)
	if err != nil {
		return errors.New("Shinefile does not contain valid json")
	}

	// Setup termui
	if err := ui.Init(); err != nil {
		return errors.New("Erorr with termui init")
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()
	dir := "root"
	var shown []string
	var ids []int

	// Create the list and populate it
	list := widgets.NewList()
	list.Title = shine.Metadata["title"]
	list.SetRect(0, 0, width, height)
	list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)

	for i, j := range shine.Steps {
		if j.Folder == dir {
			shown = append(shown, j.Text)
			ids = append(ids, i)
		}
	}

	list.Rows = shown

	ui.Render(list)

	uiEvent := ui.PollEvents()

	for {
		// Handle input
		e := <-uiEvent
		switch e.ID {
		case "<Down>", "j":
			list.ScrollDown()
		case "<Up>", "k":
			list.ScrollUp()
		case "<Enter>":
			action := shine.Steps[ids[list.SelectedRow]]
			switch action.Method {
			// Handle method cd, which changes displayed elements
			case "cd":
				dir = action.Data

				ids = nil
				shown = nil

				for i, j := range shine.Steps {
					if j.Folder == dir {
						shown = append(shown, j.Text)
						ids = append(ids, i)
					}
				}

				list.Rows = shown
			case "command", "run":
				command(action.Data)
			case "exit":
				return nil
			}
		case "q", "<C-c>":
			return nil
		}
		ui.Render(list)
	}

	return nil
}

// Ask a user a specific question
func prompt(question string) (string, error) {
	// Setup termui
	if err := ui.Init(); err != nil {
		return "", errors.New("Erorr with termui init")
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()

	// Create
	para := widgets.NewParagraph()
	para.SetRect(0, 0, width, height)
	para.Text = question
	ui.Render(para)

	var cursorFlash bool
	var editline string

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second / 8).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<C-c>":
				return "Default", nil
			case "<Backspace>":
				if len(editline) > 0 {
					editline = editline[0 : len(editline)-1]
				}
			case "<Space>":
				editline += " "
			case "<Enter>":
				return editline, nil
			default:
				if len(e.ID) == 1 {
					editline += e.ID
				}
			}

			if cursorFlash {
				para.Text = question + editline + "_"
			} else {
				para.Text = question + editline
			}

			ui.Render(para)

		case <-ticker:
			cursorFlash = !cursorFlash
			if cursorFlash {
				para.Text = question + editline + "_"
			} else {
				para.Text = question + editline
			}
			ui.Render(para)
		}
	}
}
