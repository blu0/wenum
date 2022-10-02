package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	file, _ := os.Open("./file.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	for _, cli := range lines {

		cmd := exec.Command("cmd", "/C", cli)
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		cmd.Stderr = stderr
		cmd.Stdout = stdout

		fmt.Printf("Command: %q\n", cli)

		output := cmd.Run()
		if output != nil {
			fmt.Println("Error: ", stderr.String())
		}

		fmt.Println("Output: ", stdout.String())
	}
}
