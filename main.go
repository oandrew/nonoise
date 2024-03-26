package main

import (
	"bytes"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf
	if err := cmd.Run(); err != nil {
		stderrBuf.WriteTo(os.Stderr)
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitCode := exitErr.ExitCode(); exitCode > 0 {
				os.Exit(exitCode)
			}
		}
		os.Exit(1)
	}
}
