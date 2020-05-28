package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func shellout(command string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func readFromCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("commit msg: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	// commitArgs := fmt.Sprintf("commit -m \"%s\"", text)
	// fmt.Println(commitArgs)
	return text
}

func outputError(out string, errout string, err error) {
	log.Printf("Command finished with error: %v", err)
	log.Printf("stdout: %v", out)
	log.Fatalf("stderr: %v", errout)
}

func main() {
	errout, out, err := shellout("git", "add", ".")
	if err != nil {
		outputError(out, errout, err)
	} else {
		log.Printf(out)
	}

	commitArgs := readFromCommand()
	errout, out, err = shellout("git", "commit", "-m", commitArgs)
	if err != nil {
		outputError(out, errout, err)
	} else {
		log.Printf(out)
	}

	errout, out, err = shellout("git", "push")
	if err != nil {
		outputError(out, errout, err)
	} else {
		log.Printf(out)
	}
}
