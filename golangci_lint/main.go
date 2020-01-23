package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("./bin/golangci-lint.exe", "run", "--fix")
		if err := cmd.Run(); err != nil {
			fmt.Printf("golangci-lint error: %v", err)
		}
	} else {
		cmd := exec.Command("./bin/golangci-lint", "run", "--fix")
		if err := cmd.Run(); err != nil {
			fmt.Printf("golangci-lint error: %v", err)
		}
	}
}
