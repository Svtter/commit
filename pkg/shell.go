package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Shellout function to capture command output
func Shellout(command string, args ...string) RunResult {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return RunResult{stdout.String(), stderr.String(), err}
}

// ReadFromCommand read commmit msg from command line.
func ReadFromCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("commit msg: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	// commitArgs := fmt.Sprintf("commit -m \"%s\"", text)
	// fmt.Println(commitArgs)
	return text
}
