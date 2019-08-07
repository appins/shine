package main

import (
	"os"
	"os/exec"

	ui "github.com/gizak/termui/v3"
)

func command(cmdStr string) {
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	ui.Close()
	cmd.Run()
	ui.Init()

}
